var garden;
(function () {
  const bodyWidth = document.body.offsetWidth;
  const bodyHeight = document.body.offsetHeight;
  const offsetSize = bodyWidth < bodyHeight ? bodyWidth : bodyHeight;
  const canvas = document.createElement("canvas");
  canvas.width = 500;
  canvas.height = 500;
  const $heart = document.querySelector("#heart");
  $heart.append(canvas);
  $heart.style.transform = `scale(${offsetSize / 500})`;

  const ctx = canvas.getContext("2d");
  garden = new Garden(ctx, canvas);

  startHeartAnimation();

  // renderLoop
  setInterval(function () {
    garden.render();
  }, Garden.options.growSpeed);
})();

// x最大值429.2403441681815
// y最大值456.96640482901796
const canvasWidth = 500;
function getHeartPoint(angle) {
  var t = angle / Math.PI;
  var x = 19.5 * (16 * Math.pow(Math.sin(t), 3));
  var y =
    -20 *
      (13 * Math.cos(t) -
        5 * Math.cos(2 * t) -
        2 * Math.cos(3 * t) -
        Math.cos(4 * t)) -
    50;
  var zoom = (457 * 1.6) / canvasWidth;

  return new Array(canvasWidth / 2 + x / zoom, canvasWidth / 2 + y / zoom);
}

function startHeartAnimation() {
  var angle = 10;
  var heart = new Array();

  var animationTimer = setInterval(function () {
    var draw = true;
    var bloom = getHeartPoint(angle);
    for (var i = 0; i < heart.length; i++) {
      var p = heart[i];
      var distance = Math.sqrt(
        Math.pow(p[0] - bloom[0], 2) + Math.pow(p[1] - bloom[1], 2)
      );
      if (distance < Garden.options.bloomRadius.max * 1.3) {
        draw = false;
        break;
      }
    }
    if (draw) {
      heart.push(bloom);
      garden.createRandomBloom(bloom[0], bloom[1]);
    }
    if (angle >= 30) {
      clearInterval(animationTimer);
      startWords();
    } else {
      angle += 0.2;
    }
  }, 50);
}

function startWords() {
  var together = new Date();
  together.setFullYear(2022, 3, 6);
  together.setHours(20);
  together.setMinutes(0);
  together.setSeconds(0);
  together.setMilliseconds(0);

  timeElapse(together);
  setInterval(function () {
    timeElapse(together);
  }, 500);

  document.querySelector("#messages").style.opacity = 1;
  document.querySelector("#loveu").style.opacity = 1;
}
function timeElapse(date) {
  var current = Date();
  var seconds = (Date.parse(current) - Date.parse(date)) / 1000;
  var days = Math.floor(seconds / (3600 * 24));
  seconds = seconds % (3600 * 24);
  var hours = Math.floor(seconds / 3600);
  if (hours < 10) {
    hours = "0" + hours;
  }
  seconds = seconds % 3600;
  var minutes = Math.floor(seconds / 60);
  if (minutes < 10) {
    minutes = "0" + minutes;
  }
  seconds = seconds % 60;
  if (seconds < 10) {
    seconds = "0" + seconds;
  }
  var result =
    '<span class="digit">' +
    days +
    '</span> days <span class="digit">' +
    hours +
    '</span> hours <span class="digit">' +
    minutes +
    '</span> minutes <span class="digit">' +
    seconds +
    "</span> seconds";
  document.querySelector("#elapseClock").innerHTML = result;
}
