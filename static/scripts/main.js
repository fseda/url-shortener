function copyToClipboard(text) {
  var textarea = document.createElement("textarea");
  textarea.value = text;
  document.body.appendChild(textarea);
  textarea.select();
  document.execCommand("copy");
  document.body.removeChild(textarea);
  
  var copyButton = document.querySelector(".copy-button");
  copyButton.querySelector(".copy-icon").textContent = "✅"; // Change to green tick
  copyButton.querySelector(".copy-subtitle").textContent = "Copied!";
}
