document.querySelector("#lable").addEventListener("click", () => {
  document.querySelector("#files").click()
})

const uploadInput = document.querySelector("#files");

const uid = () =>
  Date.now().toString(36) + Math.random().toString(36).substr(2);

  uploadInput.addEventListener("change", () => {
    let numberOfBytes = 0;
    for (const file of uploadInput.files) {
      file.id = uid();
      numberOfBytes += file.size;
    }

    const units = [ "B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"];

    const exponent = Math.min(
      Math.floor(Math.log(numberOfBytes) / Math.log(1024)),
      units.length - 1
    );
    const approx = numberOfBytes / 1024 ** exponent;
    const output =
      exponent === 0
        ? `${numberOfBytes} bytes`
        : `${approx.toFixed(3)} ${
            units[exponent]
          } (${numberOfBytes} bytes)`;

    for (let i = 0; i < uploadInput.files.length; i++) {
      const file = uploadInput.files[i];
      uploadFile(file);
    }
  },
  false
);

const uploadFile = async (file) => {
  initProgressBar(file);

  const req = new XMLHttpRequest();
  req.upload.addEventListener("progress", (e) => updateProgress(e, file));
  req.open("POST", "/load");

  const form = new FormData();
  form.append("file", file, file.name);

  req.onreadystatechange = function () {
    if (this.readyState == 4 && this.status == 200) {
      const data = JSON.parse(this.responseText);

      document.querySelector(`#${file.id}`).style.cursor = "pointer";
      document
        .querySelector(`#${file.id}`)
        .addEventListener("click", () => {
          navigator.clipboard.writeText(window.location.href + data.name);
        });
    }
  };

  req.send(form);
};

const updateProgress = (e, file) => {
  document.querySelector("#value" + file.id).style.width =
    Math.round((e.loaded / e.total) * 100) + "%";

  if (Math.round((e.loaded / e.total) * 100) === 100) {
    document.querySelector("#value" + file.id).style.backgroundColor =
      "#2bc253";
  }
};

function initProgressBar(file) {
  const progressItem = document.createElement("div");
  progressItem.id = file.id;

  const filename =
    file.name.length > 14
      ? file.name.slice(0, 6) +
        "..." +
        file.name.slice(file.name.length - 6, file.name.length)
      : file.name;
  const extention = file.name.split(".")[file.name.split(".").length - 1];

  let defaulImage = "/assets/document.png";

  const media = [ "MKV", "MOV", "MP4", "M4V", "MP4V", "3G2", "3GP2", "3GP", "3GPP", "M2TS", "IVF", "MPG", "MPEG", "M1V", "MP2", "MP3", "MPA", "MPE", "M3U", "WMD" ];
  const pdf = [ "MOV", "MP4", "M4V", "MP4V", "3G2", "3GP2", "3GP", "3GPP", "M2TS", "IVF", "MPG", "MPEG", "M1V", "MP2", "MP3", "MPA", "MPE", "M3U", "WMD" ];

  if (media.includes(extention.toUpperCase()))
    defaulImage = "/assets/video.png";
  if (extention == "pdf") defaulImage = "/assets/pdf.png";

  progressItem.innerHTML = `
    <div class="progress__container" >
      <img 
        class="progress__image" 
        id="image${file.id}" 
        src="${defaulImage}"
        alt="image" />
      <span class="progress__name" >${filename}</span>
      <div class="progress__value" id="value${file.id}"  style="width: 11%" ></div>
    </div>`;

  document.querySelector(".progress").appendChild(progressItem);

  const allowFormats = [
    "image/png",
    "image/svg+xml",
    "image/jpg",
    "image/jpeg",
  ];

  if (allowFormats.includes(file.type)) {
    var fr = new FileReader(file);
    fr.onload = (evt) => {
      document.querySelector(`#image${file.id}`).src = evt.target.result;
    };

    fr.readAsDataURL(file);
  }
}