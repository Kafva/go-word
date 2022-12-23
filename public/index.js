const showWordDefinition = (data,timeout) => {
    new Promise(r => setTimeout(r, timeout)).then(() => {
        fetch(`https://api.dictionaryapi.dev/api/v2/entries/en/${data}`)
            .then(res => res.json()).then( data => {
                if (data.length > 0 && 'meanings' in data[0] &&
                    data[0].meanings.length > 0 &&
                    data[0].meanings[0].definitions.length > 0) {

                    const def = data[0].meanings[0].definitions[0].definition
                    const capitalized = def.charAt(0).toUpperCase() + def.slice(1);
                    document.querySelector('h4').innerText = '"' + capitalized + '"'
                }
            })
    })
}

setTimeout(showWordDefinition, 5000);
