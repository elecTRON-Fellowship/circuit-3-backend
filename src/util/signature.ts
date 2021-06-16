import crypto from "crypto-js";

export const calcSignature = (
  httpMethod: string,
  url: string,
  salt: string,
  accessKey: string,
  secretKey: string,
  body: string,
): string => {
  const timestamp: string = (
    Math.floor(new Date().getTime() / 1000) - 10
  ).toString();
  const toSign: string =
    httpMethod + url + salt + timestamp + accessKey + secretKey + body;
	const signature = crypto.enc.Hex.stringify(crypto.HmacSHA256(toSign, secretKey))
	return crypto.enc.Base64.stringify(crypto.enc.Utf8.parse(signature))
};
