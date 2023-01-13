import { $, getData, on, post } from "./lib.js";
import { webauthnLogin, webauthnRegister } from "./webauthn.js";

on(document, "DOMContentLoaded", () => {
  let loginResponse: {totpToken: string};

  on($("#register")!, "click", async (e) => {
    e.preventDefault();
    const loginResponse = await post("./api/login", {
      email: "sandrolain@outlook.com",
      password: "test123456"
    }) as {totpToken: string};
    const totpuri = "otpauth://totp/identity:sandrolain@outlook.com?algorithm=SHA1&digits=6&issuer=identity&period=30&secret=Q7CMTR346KKBG6JST6PCWIAP4UN7CW4W";
    const totpCode = await (await fetch(`https://www.sandrolain.com/.netlify/functions/totpCode?uri=${encodeURIComponent(totpuri)}`)).text();
    const response = await post("./api/loginConfirm", {totpCode}, loginResponse.totpToken) as {sessionToken: string};

    webauthnRegister(response.sessionToken);
  });

  on($("#login")!, "click", async (e) => {
    e.preventDefault();
    const email = "sandrolain@outlook.com";
    webauthnLogin(email);
  });
});