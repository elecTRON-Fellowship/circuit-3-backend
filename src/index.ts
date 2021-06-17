import express from "express";
import json from "body-parser";
import crypto from "crypto-js";
import { calcSignature } from "./util/signature";
import axios from "axios";
import walletRouter from "./routes/create_wallet";
require("dotenv").config();

const app = express();

app.use(json());

app.get("/", (_req: express.Request, res: express.Response) => {
  res.send("Hey There");
});

app.use("/", walletRouter.router);

app.get(
  "/currencies",
  async (_req: express.Request, _res: express.Response) => {
    const accessKey: string = process.env.RAPYD_ACCESS_KEY!;
    const secretKey: string = process.env.RAPYD_SECRET_KEY!;

    const salt: string = crypto.lib.WordArray.random(12).toString();
    const timestamp: string = (
      Math.floor(new Date().getTime() / 1000) - 10
    ).toString();
    const signature = calcSignature(
      "get",
      "/v1/data/currencies",
      salt,
      accessKey,
      secretKey,
      "",
    );

    await axios
      .get("https://sandboxapi.rapyd.net/v1/data/currencies", {
        headers: {
          access_key: accessKey,
          salt: salt,
          timestamp: timestamp,
          signature: signature,
        },
      })
      .then(res => {
        _res.send(res.data);
      })
      .catch(err => {
        console.log(err);
      });
  },
);

app.listen(process.env.PORT ||9000, () => {
  console.log("server started on port 9000...");
});
