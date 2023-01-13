
export function $<T extends Element>(selector: string) {
  return document.querySelector<T>(selector);
}

export function on(el: Element | Document, name: string, fn: EventListenerOrEventListenerObject) {
  el.addEventListener(name, fn);
}

type Data = Record<string, FormDataEntryValue>;

export function getData(el: HTMLFormElement) {
  const data = new FormData(el)
  const res: Data = {};
  data.forEach((v, k) => { res[k] = v; } )
  return res
}

export async function post<T=any>(url: string, data: any, token?: string): Promise<T> {
  const headers: Record<string, string> = {
    "Content-Type": "application/json"
  };
  if (token) {
    headers["Authorization"] = `Bearer ${token}`;
  }
  const res = await fetch(url, {
    method: "POST",
    body: JSON.stringify(data),
    headers
  });
  return res.json();
}

// Base64 to ArrayBuffer
export function bufferDecode(value: string) {
  value = value
    .replace(/-/g, '+')
    .replace(/_/g, '/');
  // Pad out with standard base64 required padding characters
  var pad = value.length % 4;
  if (pad) {
    if(pad === 1) {
      throw new Error('InvalidLengthError: Input base64url string is the wrong length to determine padding');
    }
    value += new Array(5-pad).join('=');
  }
  return Uint8Array.from(atob(value), c => c.charCodeAt(0));
}
export function bufferEncode(value: ArrayBuffer) {
  return btoa(String.fromCharCode.apply(null, Array.from(new Uint8Array(value))))
    .replace(/\+/g, "-")
    .replace(/\//g, "_")
    .replace(/=/g, "");
}
