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

const payloads = createPayloads()
const hasAdvisorToken = Boolean(process.env.ADVISOR_TOKEN)

let advisor: AdvisorCore | undefined
let storageDownloadPayloadPromise: Promise<StorageDownloadPayload> | undefined

function getAdvisor(): AdvisorCore {
  if (!advisor) {
    advisor = createAdvisor()
  }

  return advisor
}

function getStorageDownloadPayload(): Promise<StorageDownloadPayload> {
  if (!storageDownloadPayloadPromise) {
    storageDownloadPayloadPromise = resolveStorageDownloadPayload(
      getAdvisor(),
      payloads.storageListPayload,
    )
  }

  return storageDownloadPayloadPromise
}

if (!hasAdvisorToken) {
  test("integration setup requires ADVISOR_TOKEN", () => {
    throw new Error("Set ADVISOR_TOKEN before running the Node integration tests.")
  })
} else {
  for (const methodName of ["getDaily", "getHourly", "getPeriod"] as const) {
    test(`forecast.${methodName}`, async () => {
      const response = await getAdvisor().forecast[methodName](payloads.weatherPayload)
      assertJsonSuccess(response)
    })
  }

  for (const methodName of ["getDaily", "getHourly", "getPeriod"] as const) {
    test(`observed.${methodName}`, async () => {
      const response = await getAdvisor().observed[methodName](payloads.weatherPayload)
      assertJsonSuccess(response)
    })
  }

  test("observed.getStationData", async () => {
    const response = await getAdvisor().observed.getStationData(payloads.stationPayload)
    assertJsonSuccess(response)
  })

  for (const methodName of ["getFireFocus", "getLightning"] as const) {
    test(`observed.${methodName}`, async () => {
      const response = await getAdvisor().observed[methodName](payloads.radiusPayload)
      assertJsonSuccess(response)
    })
  }

  for (const methodName of ["getFireFocusByGeometry", "getLightningByGeometry"] as const) {
    test(`observed.${methodName}`, async () => {
      const response = await getAdvisor().observed[methodName](payloads.geometryPayload)
      assertJsonSuccess(response)
    })
  }

  test("observed.getLightningDetails", async () => {
    const response = await getAdvisor().observed.getLightningDetails(payloads.lightningDetailsPayload)
    assertJsonSuccess(response)
  })

  test("observed.getLightningLite", async () => {
    const response = await getAdvisor().observed.getLightningLite(payloads.lightningLitePayload)
    assertJsonSuccess(response)
  })

  test("currentWeather.get", async () => {
    const response = await getAdvisor().currentWeather.get(payloads.currentWeatherPayload)
    assertJsonSuccess(response)
  })

  for (const methodName of ["getDaily", "getMonthly"] as const) {
    test(`climatology.${methodName}`, async () => {
      const response = await getAdvisor().climatology[methodName](payloads.climatologyPayload)
      assertJsonSuccess(response)
    })
  }

  test("monitoring.getAlerts", async () => {
    const response = await getAdvisor().monitoring.getAlerts()
    assertJsonSuccess(response)
  })

  test("stations.getLastData", async () => {
    const response = await getAdvisor().stations.getLastData(payloads.stationsLastDataPayload)
    assertJsonSuccess(response)
  })

  test("plan.getInfo", async () => {
    const response = await getAdvisor().plan.getInfo(payloads.planInfoPayload)
    assertJsonSuccess(response)
  })

  test("plan.getRequestDetails", async () => {
    const response = await getAdvisor().plan.getRequestDetails(payloads.requestDetailsPayload)
    assertJsonSuccess(response)
  })

  test("plan.getLocale", async () => {
    const response = await getAdvisor().plan.getLocale(payloads.planLocalePayload)
    assertJsonSuccess(response)
  })

  for (const methodName of [
    "getForecastDaily",
    "getForecastHourly",
    "getObservedDaily",
    "getObservedHourly",
  ] as const) {
    test(`chart.${methodName}`, async () => {
      const response = await getAdvisor().chart[methodName](payloads.weatherChartPayload)
      assertBufferSuccess(response)
    })
  }

  test("storage.listFiles", async () => {
    const response = await getAdvisor().storage.listFiles(payloads.storageListPayload)
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
    const response = await getAdvisor().staticMap.get(payloads.staticMapPayload)
    assertBufferSuccess(response)
  })

  test("tms.get", async () => {
    const response = await getAdvisor().tms.get(payloads.tmsPayload)
    assertBufferSuccess(response)
  })

  test("pmtiles.get", async () => {
    const response = await getAdvisor().pmtiles.get(payloads.pmtilesPayload)
    assertBufferSuccess(response)
  })

  test("schema.getDefinition", async () => {
    const response = await getAdvisor().schema.getDefinition()
    assertJsonSuccess(response)
  })

  test("schema.postDefinition", async () => {
    const response = await getAdvisor().schema.postDefinition(payloads.schemaDefinitionPayload)
    assertJsonSuccess(response)
  })
}
