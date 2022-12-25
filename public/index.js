const showWordDefinition = () => {
    h1 = document.querySelector('h1').innerText

    fetch(`https://api.dictionaryapi.dev/api/v2/entries/en/${h1}`)
        .then(res => res.json()).then(data => {
            console.log(data)
            if (data.length > 0 && 'meanings' in data[0] &&
                data[0].meanings.length > 0 &&
                data[0].meanings[0].definitions.length > 0) {

                const def = data[0].meanings[0].definitions[0].definition
                const capitalized = def.charAt(0).toUpperCase() + def.slice(1);

                const h4 = document.querySelector('h4')
                h4.innerText = '"' + capitalized + '"'
                h4.style.opacity = 1.0
            }
        })
}

setTimeout(showWordDefinition, 2000);
