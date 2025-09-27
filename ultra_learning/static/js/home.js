const startBtn = document.querySelector(".button-init")
const stopBtn = document.querySelector(".button-stop")
const timeDisplay = document.querySelector(".time span")

let startTimestamp = null
let rafId = null
let elapsed = null

function formatTime(totalSeconds) {
  const min = String(Math.floor(totalSeconds / 60)).padStart(2, "0");
  const sec = String(totalSeconds % 60).padStart(2, "0");
  return `${min}:${sec}`;
}

startBtn.addEventListener('click', e => {
    startTimestamp = Date.now()
    
    rafId = setInterval(() => {
        elapsed = Math.floor((Date.now() - startTimestamp) / 1000)
        timeDisplay.textContent = formatTime(elapsed)
    }, 1000)

})

stopBtn.addEventListener('click', e => {
    clearInterval(rafId)
    timeDisplay.textContent = "00:00"
    console.log(e.currentTarget.dataset.today)
    console.log(elapsed / 60)
})