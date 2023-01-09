import { post } from "./lib.js";

export async function webauthRegister(token: string, email: string) {

  post("./api/")

  const res = await fetch('/.netlify/functions/registerBegin?email=' + encodeURIComponent(email));
  const credentialCreationOptions = await res.json()


  credentialCreationOptions.publicKey.challenge = bufferDecode(credentialCreationOptions.publicKey.challenge);
  credentialCreationOptions.publicKey.user.id = bufferDecode(credentialCreationOptions.publicKey.user.id);

  const {publicKey} = credentialCreationOptions;

  const credential = await navigator.credentials.create({publicKey})

  const {id, type, rawId} = credential;
  const {attestationObject, clientDataJSON} = credential.response;

  const data = {
    id,
    type,
    rawId: bufferEncode(rawId),
    response: {
      attestationObject: bufferEncode(attestationObject),
      clientDataJSON: bufferEncode(clientDataJSON),
    }
  }
  console.log("🚀 ~ file: main.js ~ line 24 ~ registerUser ~ data", data)

  const finishResponse = await fetch('/.netlify/functions/registerFinish?email=' + encodeURIComponent(email), {
    method: "POST",
    body: JSON.stringify(data)
  })
  const finishData = await finishResponse.json();
  console.log("🚀 ~ file: main.js ~ line 38 ~ registerUser ~ finishData", finishData)

  alert("successfully registered !")
}