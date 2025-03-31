
// 2. Acortar URL con el token
async function shortenUrl() {
  const longUrl = document.getElementById('longUrl').value;
  const token = localStorage.getItem('token');
  const headers = { 'Content-Type': 'application/json' };

  console.log("shortenUrl Token: ", token)

  if (token) {
    headers['Authorization'] = `Bearer-${token}`;
  }

  const res = await fetch('/shorten', {
    method: 'POST',
    headers,
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
