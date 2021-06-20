import admin from "firebase-admin";
require("dotenv").config();

admin.initializeApp({
  credential: admin.credential.cert(
    JSON.parse(Buffer.from(process.env.FIREBASE!, "base64").toString("ascii")),
  ),
});

export const db = admin.firestore();
