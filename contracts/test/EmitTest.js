const EmitEventMock = artifacts.require("EmitEventMock")
const { assert } = require("chai");
const truffleAssert = require("truffle-assertions");


contract("EmitEventMock", async () => {
    it("test", async () => {
        cont = await EmitEventMock.new();
        result = await cont.test();
        console.log(JSON.stringify(result.receipt.rawLogs));
        // assert.equal(true, false);
    })
});
