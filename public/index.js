const show_def = (data,timeout) => {
  new Promise(r => setTimeout(r, timeout)).then( _ => {
    fetch(`https://api.dictionaryapi.dev/api/v2/entries/en/${data}`)
      .then(res => res.json() ).then( data => {

        if (data.length > 0 && 'meanings' in data[0] &&
          data[0].meanings.length > 0 &&
          data[0].meanings[0].definitions.length > 0) {

          let def = data[0].meanings[0].definitions[0].definition
          def     = def.charAt(0).toUpperCase() + def.slice(1);
          document.querySelector('h4').innerText = '"' + def + '"'
        }
      })
  })

}

fetch('/word', { method: 'GET', mode: 'no-cors' })
  .then(res => res.text())
  .then(word => {
    document.querySelector('h1').innerText = word ? word : ""
    document.title = word ? word : "New tab"
    // Show definition after X seconds
    // note that the 'animation-delay' on the <h4> element needs
    // to be equal to the timeout
    show_def(word, 8000)
  })
