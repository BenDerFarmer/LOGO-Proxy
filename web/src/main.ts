import "./style.css";

document.querySelector<HTMLDivElement>("#app")!.innerHTML = `
  <div>
    <h1 id="title">Merker</h1>
    <p id="value"></p>
    <input type="number" max="64" min="1" id="input" value="1"/>
    <button id="on">Ein </button>
    <button id="off">Aus</button>
 </div>
`;

function registerEvent() {
  const onBtn = document.getElementById("on");
  const offBtn = document.getElementById("off");
  const input = document.getElementById("input") as HTMLInputElement;
  const title = document.getElementById("value");
  if (onBtn == null || offBtn == null || input == null || title == null) return;

  onBtn.addEventListener("click", async () => {
    setM(input.value, "1");
  });

  offBtn.addEventListener("click", async () => {
    setM(input.value, "0");
  });

  setInterval(async () => {
    const value = await getM(input.value);
    title.innerText = value == "1" ? "on" : "off";
  }, 400);
}

function setM(id: string, value: string) {
  fetch("http://localhost:3000/m/" + id + "/" + value, {
    method: "POST",
  });
}

async function getM(id: string): Promise<string> {
  const resp = await fetch("http://localhost:3000/m/" + id);

  const text = await resp.text();
  return text;
}

registerEvent();
