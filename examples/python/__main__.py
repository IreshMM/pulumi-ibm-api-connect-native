import pulumi
import pulumi_ibm-api-connect as ibm-api-connect

my_random_resource = ibm-api-connect.Random("myRandomResource", length=24)
pulumi.export("output", {
    "value": my_random_resource.result,
})
