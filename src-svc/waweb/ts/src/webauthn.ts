import { bufferDecode, bufferEncode, post } from "./lib.js";

export async function webauthnRegister(token: string): Promise<{sessionToken: string}> {
  const credentialCreation = await post("./api/webauthn/beginRegister", null, token)
  const {publicKey} = credentialCreation;
  publicKey.challenge = bufferDecode(publicKey.challenge);
  publicKey.user.id   = bufferDecode(publicKey.user.id);
  const credential = await navigator.credentials.create({publicKey}) as PublicKeyCredential;
  if (!credential) {
    throw new Error("Unable to create new credential");
  }
  const {id, type, rawId} = credential;
  const {attestationObject, clientDataJSON} = credential.response as AuthenticatorAttestationResponse;
  const data = {
    id,
    type,
    rawId: bufferEncode(rawId),
    response: {
      attestationObject: bufferEncode(attestationObject),
      clientDataJSON: bufferEncode(clientDataJSON),
    }
  };
  return await post("./api/webauthn/finishRegister", data, token) as {sessionToken: string};
}

export async function webauthnLogin(email: string): Promise<{sessionToken: string}> {
  const {credentialAssertion, webauthnToken} = await post("./api/webauthn/beginLogin", {email})
  const {publicKey} = credentialAssertion;
  publicKey.challenge = bufferDecode(publicKey.challenge);
  publicKey.allowCredentials.forEach(function (listItem: any) {
    listItem.id = bufferDecode(listItem.id)
  });
  const credential = await navigator.credentials.get({publicKey}) as PublicKeyCredential;
  if (!credential) {
    throw new Error("Unable to request credential");
  }
  const {id, type, rawId} = credential;
  const {authenticatorData, clientDataJSON, signature, userHandle} = credential.response as AuthenticatorAssertionResponse;
  const data = {
    id,
    type,
    rawId: bufferEncode(rawId),
    response: {
      authenticatorData: bufferEncode(authenticatorData),
      clientDataJSON: bufferEncode(clientDataJSON),
      signature: bufferEncode(signature),
      userHandle: userHandle ? bufferEncode(userHandle) : null
    }
  };
  return await post("./api/webauthn/finishLogin", data, webauthnToken) as {sessionToken: string};
}