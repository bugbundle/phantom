<!DOCTYPE html>
<html lang="en" data-theme="pastel">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Phantom dashboard</title>
  <script defer src="https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js"></script>
</head>
<body>
  <nav>
    <a>Phantom</a>
  </nav>
  <main>
    <div
      x-init="status = await (await fetch('/cameras/status')).body"
      x-data="{
        status: false,
        async toggleCamera() {
          const response = await fetch('/cameras', {method: status ? 'DELETE' : 'POST'});
          const responseStatus = await response.status;
          status = !status;
        }
      }">
      <img src="/cameras" x-show="status" style='width: 640px; height: 360px;'/>
      <div>
        <button
          x-text="status ? 'stop' : 'start'"
          x-on:click="toggleCamera()"/>
      </div>
    </div>
  </main>
</body>
</html>
