using System.Globalization;
using StormGeo.AdvisorCore;
using StormGeo.AdvisorCore.Payloads;
using Xunit;

namespace StormGeo.AdvisorCore.Tests.Integration;

public sealed class IntegrationFactAttribute : FactAttribute
{
    public IntegrationFactAttribute()
    {
        if (string.IsNullOrWhiteSpace(Environment.GetEnvironmentVariable("ADVISOR_TOKEN")))
        {
            Skip = "Set ADVISOR_TOKEN before running the .NET integration tests.";
        }
    }
}

public sealed class IntegrationPayloads
{
    public required WeatherPayload WeatherPayload { get; init; }
    public required WeatherPayload WeatherChartPayload { get; init; }
    public required ClimatologyPayload ClimatologyPayload { get; init; }
    public required CurrentWeatherPayload CurrentWeatherPayload { get; init; }
    public required StationPayload StationPayload { get; init; }
    public required RadiusPayload RadiusPayload { get; init; }
    public required GeometryPayload GeometryPayload { get; init; }
    public required RadiusPayload LightningDetailsPayload { get; init; }
    public required LightningLitePayload LightningLitePayload { get; init; }
    public required StorageListPayload StorageListPayload { get; init; }
    public required PlanInfoPayload PlanInfoPayload { get; init; }
    public required RequestDetailsPayload RequestDetailsPayload { get; init; }
    public required StaticMapPayload StaticMapPayload { get; init; }
    public required TmsPayload TmsPayload { get; init; }
    public required PmtilesPayload PmtilesPayload { get; init; }
    public required string SchemaIdentifier { get; init; }
    public required Dictionary<string, SchameDefinitionField> SchemaDefinitionFields { get; init; }
    public required Dictionary<string, object> SchemaParameters { get; init; }
}

public static class IntegrationHelpers
{
    private const string DefaultStationId = "bWV0b3M6MDM0MTMyRjM6LTIyLjIzMTQ1MjQ4MDg0NDU2Oi00NC4yNTEzNTMwMzgzMTcx";
    private const string DefaultGeometry = "{\"type\":\"Polygon\",\"coordinates\":[[[-47.09861059094109,-23.280351816702165],[-47.09861059094109,-23.895097240590488],[-46.12890390857018,-23.895097240590488],[-46.12890390857018,-23.280351816702165],[-47.09861059094109,-23.280351816702165]]]}";
    private const string DefaultStorageFileName = "Boletim_Meteorologico_-_Andre_Alves-22_03_2026_07_02.pdf";
    private const string DefaultStorageAccessKey = "ad441946268c2227a96ad254fafe71565acd02e2-1774173734244";

    public static AdvisorCore CreateSdk()
    {
        var token = Environment.GetEnvironmentVariable("ADVISOR_TOKEN");
        if (string.IsNullOrWhiteSpace(token))
        {
            throw new InvalidOperationException("Set ADVISOR_TOKEN before running the .NET integration tests.");
        }

        var sdk = new AdvisorCore(token, attempts: 1, delayInSeconds: 0);
        sdk.SetHeaderAccept("application/json");
        sdk.SetHeaderAcceptLanguage(Environment.GetEnvironmentVariable("ADVISOR_ACCEPT_LANGUAGE") ?? "en-US");
        return sdk;
    }

    public static IntegrationPayloads CreatePayloads()
    {
        var localeId = GetEnvInt("ADVISOR_LOCALE_ID", 3477);
        var stationId = Environment.GetEnvironmentVariable("ADVISOR_STATION_ID") ?? DefaultStationId;
        var today = DateTime.Now;
        var observedDay = today.AddDays(-1);
        var observedPeriodEnd = observedDay;
        var observedPeriodStart = observedPeriodEnd.AddDays(-4);
        var forecastDay = today.AddDays(1);
        var forecastHourEnd = new DateTime(forecastDay.Year, forecastDay.Month, forecastDay.Day, 1, 0, 0, forecastDay.Kind);
        const string schemaIdentifier = "schemaIdentifier";

        return new IntegrationPayloads
        {
            WeatherPayload = new WeatherPayload
            {
                LocaleId = localeId,
                Variables = ["temperature"],
            },
            WeatherChartPayload = new WeatherPayload
            {
                LocaleId = localeId,
                Variables = ["temperature", "precipitation"],
            },
            ClimatologyPayload = new ClimatologyPayload
            {
                LocaleId = localeId,
                Variables = ["temperature"],
            },
            CurrentWeatherPayload = new CurrentWeatherPayload
            {
                LocaleId = localeId,
            },
            StationPayload = new StationPayload
            {
                StationId = stationId,
            },
            RadiusPayload = new RadiusPayload
            {
                LocaleId = localeId,
                Radius = 10000,
            },
            GeometryPayload = new GeometryPayload
            {
                Geometry = DefaultGeometry,
                StartDate = StartOfDay(observedDay),
                EndDate = EndOfDay(observedDay),
                Radius = 10000,
            },
            LightningDetailsPayload = new RadiusPayload
            {
                Latitude = "-22.9",
                Longitude = "-43.2",
                StartDate = StartOfDay(observedDay),
                EndDate = EndOfDay(observedDay),
                Radius = 10000,
            },
            LightningLitePayload = new LightningLitePayload
            {
                Geometry = DefaultGeometry,
                StartDate = StartOfDay(observedPeriodStart),
                EndDate = EndOfDay(observedPeriodEnd),
                Radius = 10000,
                Page = 1,
                PageSize = 50,
            },
            StorageListPayload = new StorageListPayload(1, 10),
            PlanInfoPayload = new PlanInfoPayload
            {
                Timezone = -3,
            },
            RequestDetailsPayload = new RequestDetailsPayload(1, 3),
            StaticMapPayload = new StaticMapPayload
            {
                Type = "periods",
                Category = "observed",
                Variable = "temperature",
                Aggregation = "max",
                StartDate = StartOfDay(observedPeriodStart),
                EndDate = EndOfDay(observedPeriodEnd),
                Dpi = 50,
                Title = true,
                TitleVariable = "Static Map",
            },
            TmsPayload = new TmsPayload
            {
                Server = 'a',
                Mode = "forecast",
                Variable = "precipitation",
                Aggregation = "sum",
                X = 5,
                Y = 8,
                Z = 4,
                Istep = StartOfDay(forecastDay),
                Fstep = EndOfDay(forecastDay),
            },
            PmtilesPayload = new PmtilesPayload
            {
                Mode = "forecast",
                Model = "ct2w15_as",
                Variable = "precipitation",
                Aggregation = "sum",
                Istep = StartOfDay(forecastDay),
                Fstep = FormatDateTime(forecastHourEnd),
                MaxZoom = 4,
            },
            SchemaIdentifier = schemaIdentifier,
            SchemaDefinitionFields = new Dictionary<string, SchameDefinitionField>
            {
                ["arbitraryField1"] = new SchameDefinitionField
                {
                    Type = "boolean",
                    Required = true,
                    Length = 125,
                },
                ["arbitraryField2"] = new SchameDefinitionField
                {
                    Type = "number",
                    Required = true,
                },
                ["arbitraryField3"] = new SchameDefinitionField
                {
                    Type = "string",
                    Required = false,
                },
            },
            SchemaParameters = new Dictionary<string, object>
            {
                ["arbitraryField1"] = true,
                ["arbitraryField2"] = 15,
            },
        };
    }

    public static StorageDownloadPayload ResolveStorageDownloadPayload()
    {
        var fileName = Environment.GetEnvironmentVariable("ADVISOR_STORAGE_FILE_NAME");
        var accessKey = Environment.GetEnvironmentVariable("ADVISOR_STORAGE_ACCESS_KEY");

        if (!string.IsNullOrWhiteSpace(fileName) || !string.IsNullOrWhiteSpace(accessKey))
        {
            Assert.False(string.IsNullOrWhiteSpace(fileName) || string.IsNullOrWhiteSpace(accessKey),
                "Set both ADVISOR_STORAGE_FILE_NAME and ADVISOR_STORAGE_ACCESS_KEY, or neither.");

            return new StorageDownloadPayload
            {
                FileName = fileName,
                AccessKey = accessKey,
            };
        }

        return new StorageDownloadPayload
        {
            FileName = DefaultStorageFileName,
            AccessKey = DefaultStorageAccessKey,
        };
    }

    public static void AssertJsonSuccess(AdvisorResponse<string> response)
    {
        Assert.Null(response.Error);
        Assert.False(string.IsNullOrWhiteSpace(response.Data));
    }

    public static async Task AssertStreamSuccessAsync(AdvisorResponse<Stream> response)
    {
        Assert.Null(response.Error);
        Assert.NotNull(response.Data);

        await using var stream = response.Data!;
        using var memory = new MemoryStream();
        await stream.CopyToAsync(memory);
        Assert.NotEmpty(memory.ToArray());
    }

    private static int GetEnvInt(string name, int defaultValue)
    {
        var value = Environment.GetEnvironmentVariable(name);
        return int.TryParse(value, out var parsed) ? parsed : defaultValue;
    }

    private static string StartOfDay(DateTime value)
    {
        return FormatDateTime(new DateTime(value.Year, value.Month, value.Day, 0, 0, 0, value.Kind));
    }

    private static string EndOfDay(DateTime value)
    {
        return FormatDateTime(new DateTime(value.Year, value.Month, value.Day, 23, 59, 59, value.Kind));
    }

    private static string FormatDateTime(DateTime value)
    {
        return value.ToString("yyyy-MM-dd HH:mm:ss", CultureInfo.InvariantCulture);
    }
}
