import assert from "node:assert/strict"
import { existsSync, readFileSync } from "node:fs"
import path from "node:path"
import { Readable } from "node:stream"

import { AdvisorCore } from "../src/AdvisorCore"
import {
  ApiFileResponse,
  ApiResponse,
  ApiStreamResponse,
  ClimatologyPayload,
  CurrentWeatherPayload,
  GeometryPayload,
  PlanInfoPayload,
  PlanLocalePayload,
  PmtilesPayload,
  RadiusPayload,
  RequestDetailsPayload,
  StationPayload,
  StationsLastDataPayload,
  StaticMapPayload,
  StorageDownloadPayload,
  StorageListPayload,
  TmsPayload,
  WeatherPayload,
} from "../src/payloads"

type JsonLike = Record<string, unknown>

export interface IntegrationPayloads {
  weatherPayload: WeatherPayload
  weatherChartPayload: WeatherPayload
  climatologyPayload: ClimatologyPayload
  currentWeatherPayload: CurrentWeatherPayload
  stationPayload: StationPayload
  stationsLastDataPayload: StationsLastDataPayload
  radiusPayload: RadiusPayload
  geometryPayload: GeometryPayload
  lightningDetailsPayload: RadiusPayload
  lightningLitePayload: GeometryPayload & { page: number; pageSize: number }
  storageListPayload: StorageListPayload
  planInfoPayload: PlanInfoPayload
  planLocalePayload: PlanLocalePayload
  requestDetailsPayload: RequestDetailsPayload
  staticMapPayload: StaticMapPayload
  tmsPayload: TmsPayload
  pmtilesPayload: PmtilesPayload
  schemaDefinitionPayload: JsonLike
}

function getEnvInt(name: string, defaultValue: number): number {
  const value = process.env[name]
  return value ? Number.parseInt(value, 10) : defaultValue
}

function findIntegrationEnvFile(startDir: string): string | undefined {
  let currentDir = startDir

  while (true) {
    const candidate = path.join(currentDir, ".env.integration.local")
    if (existsSync(candidate)) {
      return candidate
    }

    const parentDir = path.dirname(currentDir)
    if (parentDir === currentDir) {
      return undefined
    }

    currentDir = parentDir
  }
}

function loadIntegrationEnv(): void {
  const envFile = findIntegrationEnvFile(__dirname) ?? findIntegrationEnvFile(process.cwd())
  if (!envFile) {
    return
  }

  for (const rawLine of readFileSync(envFile, "utf8").split(/\r?\n/u)) {
    const line = rawLine.trim()
    if (!line || line.startsWith("#")) {
      continue
    }

    const separatorIndex = line.indexOf("=")
    if (separatorIndex === -1) {
      continue
    }

    const name = line.slice(0, separatorIndex).trim()
    if (!name || process.env[name]) {
      continue
    }

    let value = line.slice(separatorIndex + 1).trim()
    if (
      value.length >= 2
      && ((value.startsWith("\"") && value.endsWith("\"")) || (value.startsWith("'") && value.endsWith("'")))
    ) {
      value = value.slice(1, -1)
    }

    process.env[name] = value
  }
}

function requireEnv(name: string): string {
  const value = process.env[name]
  if (!value) {
    throw new Error(
      `Set ${name} or add it to .env.integration.local before running the Node integration tests.`,
    )
  }
  return value
}

loadIntegrationEnv()

function pad(value: number): string {
  return String(value).padStart(2, "0")
}

function formatDateTime(value: Date): string {
  return [
    value.getFullYear(),
    pad(value.getMonth() + 1),
    pad(value.getDate()),
  ].join("-") + " " + [
    pad(value.getHours()),
    pad(value.getMinutes()),
    pad(value.getSeconds()),
  ].join(":")
}

function startOfDay(value: Date): string {
  return formatDateTime(new Date(
    value.getFullYear(),
    value.getMonth(),
    value.getDate(),
    0,
    0,
    0,
  ))
}

function endOfDay(value: Date): string {
  return formatDateTime(new Date(
    value.getFullYear(),
    value.getMonth(),
    value.getDate(),
    23,
    59,
    59,
  ))
}

function shiftDays(value: Date, amount: number): Date {
  const shifted = new Date(value)
  shifted.setDate(shifted.getDate() + amount)
  return shifted
}

export function createAdvisor(): AdvisorCore {
  const advisor = new AdvisorCore({
    token: requireEnv("ADVISOR_TOKEN"),
    retries: 1,
    delay: 0,
  })

  advisor.setHeaderAccept("application/json")
  advisor.setHeaderAcceptLanguage(process.env.ADVISOR_ACCEPT_LANGUAGE ?? "en-US")

  return advisor
}

export function createPayloads(): IntegrationPayloads {
  const localeId = getEnvInt("ADVISOR_LOCALE_ID", 3477)
  const planLocaleId = getEnvInt("ADVISOR_PLAN_LOCALE_ID", 5959)
  const stationId = requireEnv("ADVISOR_STATION_ID")
  const geometry = requireEnv("ADVISOR_GEOMETRY")
  const today = new Date()
  const observedDay = shiftDays(today, -1)
  const observedPeriodEnd = observedDay
  const observedPeriodStart = shiftDays(observedPeriodEnd, -4)
  const forecastDay = shiftDays(today, 1)
  const forecastHourEnd = new Date(
    forecastDay.getFullYear(),
    forecastDay.getMonth(),
    forecastDay.getDate(),
    1,
    0,
    0,
  )
  const schemaIdentifier = "schemaIdentifier"

  return {
    weatherPayload: {
      localeId,
      variables: ["temperature"],
    },
    weatherChartPayload: {
      localeId,
      variables: ["temperature", "precipitation"],
    },
    climatologyPayload: {
      localeId,
      variables: ["temperature"],
    },
    currentWeatherPayload: {
      localeId,
    },
    stationPayload: {
      stationId,
    },
    stationsLastDataPayload: {
      stationIds: [stationId],
      variables: ["temperature"],
    },
    radiusPayload: {
      localeId,
      radius: 10000,
    },
    geometryPayload: {
      geometry,
      startDate: startOfDay(observedDay),
      endDate: endOfDay(observedDay),
      radius: 10000,
    },
    lightningDetailsPayload: {
      latitude: -22.9,
      longitude: -43.2,
      startDate: startOfDay(observedDay),
      endDate: endOfDay(observedDay),
      radius: 10000,
    },
    lightningLitePayload: {
      geometry,
      startDate: startOfDay(observedPeriodStart),
      endDate: endOfDay(observedPeriodEnd),
      radius: 10000,
      page: 1,
      pageSize: 50,
    },
    storageListPayload: {
      page: 1,
      pageSize: 10,
    },
    planInfoPayload: {
      timezone: -3,
    },
    planLocalePayload: {
      localeId: planLocaleId,
    },
    requestDetailsPayload: {
      page: 1,
      pageSize: 3,
    },
    staticMapPayload: {
      type: "periods",
      category: "observed",
      variable: "temperature",
      aggregation: "max",
      startDate: startOfDay(observedPeriodStart),
      endDate: endOfDay(observedPeriodEnd),
      dpi: 50,
      title: true,
      titlevariable: "Static Map",
    },
    tmsPayload: {
      istep: startOfDay(forecastDay),
      fstep: endOfDay(forecastDay),
      server: "a",
      mode: "forecast",
      variable: "precipitation",
      aggregation: "sum",
      x: "5",
      y: "8",
      z: "4",
    },
    pmtilesPayload: {
      mode: "forecast",
      model: "ct2w15_as",
      variable: "precipitation",
      aggregation: "sum",
      istep: startOfDay(forecastDay),
      fstep: formatDateTime(forecastHourEnd),
      maxZoom: 4,
    },
    schemaDefinitionPayload: {
      identifier: schemaIdentifier,
      arbitraryField1: {
        type: "boolean",
        required: true,
        length: 125,
      },
      arbitraryField2: {
        type: "number",
        required: true,
      },
      arbitraryField3: {
        type: "string",
        required: false,
      },
    }
  }
}

export async function resolveStorageDownloadPayload(
  _advisor: AdvisorCore,
  _storageListPayload: StorageListPayload,
): Promise<StorageDownloadPayload> {
  return {
    fileName: requireEnv("ADVISOR_STORAGE_FILE_NAME"),
    accessKey: requireEnv("ADVISOR_STORAGE_ACCESS_KEY"),
  }
}

function formatError(error: unknown): string {
  if (!error) {
    return "Unknown error"
  }

  if (typeof error === "string") {
    return error
  }

  if (error instanceof Error) {
    return error.message
  }

  return JSON.stringify(error)
}

export function assertJsonSuccess(response: ApiResponse): void {
  assert.equal(response.error, null, formatError(response.error))
  assert.notEqual(response.data, null)
}

export function assertBufferSuccess(response: ApiFileResponse): void {
  assert.equal(response.error, null, formatError(response.error))
  assert.ok(Buffer.isBuffer(response.data))
  assert.ok(response.data.length > 0)
}

export async function assertStreamSuccess(response: ApiStreamResponse): Promise<void> {
  assert.equal(response.error, null, formatError(response.error))
  assert.ok(response.data)

  const chunks: Buffer[] = []
  for await (const chunk of response.data as Readable) {
    chunks.push(Buffer.isBuffer(chunk) ? chunk : Buffer.from(chunk))
  }

  assert.ok(chunks.length > 0)
  assert.ok(Buffer.concat(chunks).length > 0)
}
