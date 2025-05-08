<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Phantom dashboard</title>
  <link rel="stylesheet" href="/static/css/style.css"></link>
</head>
<body>
  <nav>
    <a>Phantom</a>
  </nav>
  <main>
      <img src="cameras" style='width: 640px; height: 360px;'/>
      <div>
        <button onclick="createCamera()">start</button>
        <button onclick="deleteCamera()">stop</button>
      </div>
  </main>
  <script>
    async function createCamera() {
      await fetch('/cameras', {method: 'POST'});
      location.reload();
    };
    async function deleteCamera() {
      await fetch('/cameras', {method: 'DELETE'});
      location.reload();
    };
  </script>
</body>
</html>
