import express from "express";
import wallet from "./routes/create_wallet";
import customer from "./routes/create_customer";
import checkout from "./routes/create_checkout";
import transfer from "./routes/transfer_funds";
import balance from "./routes/get_balance";
import card from "./routes/card_payment";
import walletFunds from "./routes/add_funds";
import listPayMethods from "./routes/list_payout_methods";
import payToCard from "./routes/pay_to_card";
import { json } from "body-parser";
import morgan from "morgan";

require("dotenv").config();

const app = express();

app.use(json());
app.use(morgan("tiny"));

app.get("/", (_req: express.Request, res: express.Response) => {
  res.send("You may close this window now...");
});
app.post("/hooks", (_req: express.Request, _res: express.Response) => {
  console.log(_req.body);
});

app.use("/", wallet.router);
app.use("/", customer.router);
app.use("/", checkout.router);
app.use("/", walletFunds.router);
app.use("/", transfer.router);
app.use("/", balance.router);
app.use("/", card.router);
app.use("/", listPayMethods.router);
app.use("/", payToCard.router);

app.listen(process.env.PORT || 9000, () => {
  console.log("server started on port 9000...");
});
