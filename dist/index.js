"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const app = express_1.default();
app.get("/", (res) => {
    res.send("Send");
});
app.listen(4200, () => {
    "listening on port 4200";
});
//# sourceMappingURL=index.js.map