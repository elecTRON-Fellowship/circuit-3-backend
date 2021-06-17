import express from "express";
import wallet from "./routes/create_wallet";
import customer from "./routes/create_customer";
import checkout from "./routes/create_checkout";
import transfer from "./routes/create_transfer";
import balance from "./routes/retreive_balance";
import walletFunds from "./routes/add_funds";
import { json } from "body-parser";
import morgan from "morgan";

require("dotenv").config();


const app = express();

app.use(json());
app.use(morgan("tiny"));

app.get("/", (_req: express.Request, res: express.Response) => {
  res.send("Hey There");
});

app.use("/", wallet.router);
app.use("/", customer.router);
app.use("/", checkout.router);
app.use("/", walletFunds.router);
app.use("/", transfer.router);
app.use("/", balance.router);

app.listen(9000, () => {
  console.log("server started on port 9000...");
});
