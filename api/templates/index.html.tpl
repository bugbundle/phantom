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
    <div>
      <img src="/cameras" style='width: 640px; height: 360px;'/>
      <div>
        <button x-data="{
          async createCamera() {
              const response = await fetch('/cameras', {method: 'POST'});
              const responseStatus = await response.status;
            }
        }" x-on:click="createCamera()">Start</button>
        <button x-data="{
          async deleteCamera() {
              const response = await fetch('/cameras', {method: 'DELETE'});
              const responseStatus = await response.status;
            }
        }" x-on:click="deleteCamera()">Stop</button>
      </div>
    </div>
  </main>
</body>
</html>
