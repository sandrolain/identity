
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

export async function post<T=any>(url: string, data: Data, token?: string): Promise<T> {
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