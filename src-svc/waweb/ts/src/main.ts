import { $, getData, on, post } from "./lib.js";

on(document, "DOMContentLoaded", () => {
  let loginResponse: {totpToken: string};

  const $login = $<HTMLFormElement>("#login")!;
  const $loginConfirm = $<HTMLFormElement>("#login-confirm")!;

  on($login, "submit", async (e) => {
    e.preventDefault();
    const data = getData($login);
    loginResponse = await post("./api/login", data);
    console.log("🚀 ~ file: login.ts:9 ~ on ~ response", loginResponse)
    $login.hidden = true;
    $loginConfirm.hidden = false;
  });

  on($loginConfirm, "submit", async (e) => {
    e.preventDefault();
    const data = getData($loginConfirm);
    const response = await post("./api/loginConfirm", data, loginResponse.totpToken);
    console.log("🚀 ~ file: login.ts:9 ~ on ~ response", response)
  });
});