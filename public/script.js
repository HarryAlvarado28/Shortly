async function shortenUrl() {
  const longUrl = document.getElementById('longUrl').value;
  const res = await fetch('/shorten', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ url: longUrl }),
  });

  const data = await res.json();
  const shortDiv = document.getElementById('shortUrl');

  if (res.ok) {
    shortDiv.innerHTML = `
        <p><strong>URL corta:</strong></p>
        <a href="${data.short_url}" target="_blank">${data.short_url}</a>
      `;
  } else {
    shortDiv.innerHTML = `<p style="color:red;">Ocurrió un error. Verifica la URL.</p>`;
  }
}
