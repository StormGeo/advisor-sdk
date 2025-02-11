namespace StormGeo.AdvisorCore.Payloads;

public class SchameDefinitionField
{
    public required string Type { get; set; }
    public required bool Required { get; set; }
    public int? Length { get; set; }

    public Dictionary<string, dynamic> ToDictionary()
    {
        var dict = new Dictionary<string, dynamic>()
        {
            { "type", Type },
            { "required", Required }
        };

        if (Length != null)
        {
            dict.Add("length", Length);
        }

        return dict;
    }
}
