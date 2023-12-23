using System.Collections.Generic;
using System.Linq;
using Pulumi;
using IbmApiConnect = Pulumi.IbmApiConnect;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new IbmApiConnect.Random("myRandomResource", new()
    {
        Length = 24,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myRandomResource.Result },
        },
    };
});

