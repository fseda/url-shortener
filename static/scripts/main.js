const copyButton = document.querySelector('.copy-button');
copyButton.addEventListener('click', () => {
  const shortenedUrl = copyButton.closest('.result').querySelector('.shortened-url a');
  if (shortenedUrl) {
    const tempInput = document.createElement('input');
    tempInput.value = shortenedUrl.innerHTML;
    document.body.appendChild(tempInput);
    tempInput.select();
    document.execCommand('copy');
    document.body.removeChild(tempInput);
    copyButton.querySelector('.copy-icon').textContent = '✅';
    }
});
