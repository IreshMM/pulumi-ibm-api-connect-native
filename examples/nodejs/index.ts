import * as pulumi from "@pulumi/pulumi";
import * as ibm-api-connect from "@pulumi/ibm-api-connect";

const myRandomResource = new ibm-api-connect.Random("myRandomResource", {length: 24});
export const output = {
    value: myRandomResource.result,
};
