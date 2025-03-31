window.addEventListener('DOMContentLoaded', async () => {
  const token = localStorage.getItem('token');

  if (!token) {
    try {
      const res = await fetch('/anon', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      const data = await res.json();

      if (res.ok && data.token) {
        localStorage.setItem('token', data.token);
        console.log('[shortly] Sesión anónima iniciada');
        document.getElementById('anonWarning').style.display = 'block';
      } else {
        console.error('Respuesta inesperada de /anon:', data);
      }
    } catch (error) {
      console.error('No se pudo iniciar sesión anónima:', error);
    }
  } else {
    // Si ya tiene sesión
    document.getElementById('anonWarning').style.display = 'none';
  }
});

async function shortenUrl() {
  const longUrl = document.getElementById('longUrl').value;
  const token = localStorage.getItem('token');
  const headers = { 'Content-Type': 'application/json' };

  console.log('shortenUrl Token: ', token);

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
