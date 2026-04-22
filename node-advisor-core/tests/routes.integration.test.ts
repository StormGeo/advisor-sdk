import { test } from "node:test"

import { AdvisorCore } from "../src/AdvisorCore"
import { StorageDownloadPayload } from "../src/payloads"
import {
  assertBufferSuccess,
  assertJsonSuccess,
  assertStreamSuccess,
  createAdvisor,
  createPayloads,
  resolveStorageDownloadPayload,
} from "./integration.helpers"

const REQUIRED_ENV_VARS = [
  "ADVISOR_TOKEN",
  "ADVISOR_STATION_ID",
  "ADVISOR_GEOMETRY",
  "ADVISOR_STORAGE_FILE_NAME",
  "ADVISOR_STORAGE_ACCESS_KEY",
] as const

const missingRequiredEnv = REQUIRED_ENV_VARS.filter((name) => !process.env[name])

let advisor: AdvisorCore | undefined
let storageDownloadPayloadPromise: Promise<StorageDownloadPayload> | undefined
let payloads: ReturnType<typeof createPayloads> | undefined

function getAdvisor(): AdvisorCore {
  if (!advisor) {
    advisor = createAdvisor()
  }

  return advisor
}

function getPayloads() {
  if (!payloads) {
    payloads = createPayloads()
  }

  return payloads
}

function getStorageDownloadPayload(): Promise<StorageDownloadPayload> {
  if (!storageDownloadPayloadPromise) {
    storageDownloadPayloadPromise = resolveStorageDownloadPayload(
      getAdvisor(),
      getPayloads().storageListPayload,
    )
  }

  return storageDownloadPayloadPromise
}

if (missingRequiredEnv.length > 0) {
  test("integration setup requires shared Advisor env", () => {
    throw new Error(
      `Set ${missingRequiredEnv.join(", ")} or add them to .env.integration.local before running the Node integration tests.`,
    )
  })
} else {
  for (const methodName of ["getDaily", "getHourly", "getPeriod"] as const) {
    test(`forecast.${methodName}`, async () => {
      const response = await getAdvisor().forecast[methodName](getPayloads().weatherPayload)
      assertJsonSuccess(response)
    })
  }

  for (const methodName of ["getDaily", "getHourly", "getPeriod"] as const) {
    test(`observed.${methodName}`, async () => {
      const response = await getAdvisor().observed[methodName](getPayloads().weatherPayload)
      assertJsonSuccess(response)
    })
  }

  test("observed.getStationData", async () => {
    const response = await getAdvisor().observed.getStationData(getPayloads().stationPayload)
    assertJsonSuccess(response)
  })

  for (const methodName of ["getFireFocus", "getLightning"] as const) {
    test(`observed.${methodName}`, async () => {
      const response = await getAdvisor().observed[methodName](getPayloads().radiusPayload)
      assertJsonSuccess(response)
    })
  }

  for (const methodName of ["getFireFocusByGeometry", "getLightningByGeometry"] as const) {
    test(`observed.${methodName}`, async () => {
      const response = await getAdvisor().observed[methodName](getPayloads().geometryPayload)
      assertJsonSuccess(response)
    })
  }

  test("observed.getLightningDetails", async () => {
    const response = await getAdvisor().observed.getLightningDetails(getPayloads().lightningDetailsPayload)
    assertJsonSuccess(response)
  })

  test("observed.getLightningLite", async () => {
    const response = await getAdvisor().observed.getLightningLite(getPayloads().lightningLitePayload)
    assertJsonSuccess(response)
  })

  test("currentWeather.get", async () => {
    const response = await getAdvisor().currentWeather.get(getPayloads().currentWeatherPayload)
    assertJsonSuccess(response)
  })

  for (const methodName of ["getDaily", "getMonthly"] as const) {
    test(`climatology.${methodName}`, async () => {
      const response = await getAdvisor().climatology[methodName](getPayloads().climatologyPayload)
      assertJsonSuccess(response)
    })
  }

  test("monitoring.getAlerts", async () => {
    const response = await getAdvisor().monitoring.getAlerts()
    assertJsonSuccess(response)
  })

  test("stations.getLastData", async () => {
    const response = await getAdvisor().stations.getLastData(getPayloads().stationsLastDataPayload)
    assertJsonSuccess(response)
  })

  test("plan.getInfo", async () => {
    const response = await getAdvisor().plan.getInfo(getPayloads().planInfoPayload)
    assertJsonSuccess(response)
  })

  test("plan.getRequestDetails", async () => {
    const response = await getAdvisor().plan.getRequestDetails(getPayloads().requestDetailsPayload)
    assertJsonSuccess(response)
  })

  test("plan.getLocale", async () => {
    const response = await getAdvisor().plan.getLocale(getPayloads().planLocalePayload)
    assertJsonSuccess(response)
  })

  for (const methodName of [
    "getForecastDaily",
    "getForecastHourly",
    "getObservedDaily",
    "getObservedHourly",
  ] as const) {
    test(`chart.${methodName}`, async () => {
      const response = await getAdvisor().chart[methodName](getPayloads().weatherChartPayload)
      assertBufferSuccess(response)
    })
  }

  test("storage.listFiles", async () => {
    const response = await getAdvisor().storage.listFiles(getPayloads().storageListPayload)
    assertJsonSuccess(response)
  })

  test("storage.downloadFile", async () => {
    const response = await getAdvisor().storage.downloadFile(await getStorageDownloadPayload())
    assertBufferSuccess(response)
  })

  test("storage.downloadFileByStream", async () => {
    const response = await getAdvisor().storage.downloadFileByStream(await getStorageDownloadPayload())
    await assertStreamSuccess(response)
  })

  test("staticMap.get", async () => {
    const response = await getAdvisor().staticMap.get(getPayloads().staticMapPayload)
    assertBufferSuccess(response)
  })

  test("tms.get", async () => {
    const response = await getAdvisor().tms.get(getPayloads().tmsPayload)
    assertBufferSuccess(response)
  })

  test("pmtiles.get", async () => {
    const response = await getAdvisor().pmtiles.get(getPayloads().pmtilesPayload)
    assertBufferSuccess(response)
  })

  test("schema.getDefinition", async () => {
    const response = await getAdvisor().schema.getDefinition()
    assertJsonSuccess(response)
  })

  test("schema.postDefinition", async () => {
    const response = await getAdvisor().schema.postDefinition(getPayloads().schemaDefinitionPayload)
    assertJsonSuccess(response)
  })
}
