import { useState } from "react"

export default function Autocomplete() {
    const [text, setText] = useState("")
    const [options, setOptions] = useState(["howdy", "partner", "how", "are", "you", "doing"])
    const [debounce, setDebounce] = useState(null)

    const onType = (e) => {
        setText(e.target.value)

        if (debounce) {
            clearTimeout(debounce)
        }

        setDebounce(setTimeout(async () => {
            // TODO(dk): call api
            fetch(`/api/autocomplete?value=${text}`).then((x) => {
                x.json().then((opts) => {
                    if (opts.data) {
                        const res_data = opts.data.map((x) => {
                            return x.name
                        })
                        setOptions([...res_data])
                    }
                }).catch((_) => {})
            }).catch((_) => {})
        }, 100))
    }

    return (
        <div className="w-full flex-col">
            <input className="border-2 py-2 px-3 w-full" type="text" onInput={onType} value={text}></input>
            { options.length > 0 && text != "" && <div className="relative shadow-md w-full">
            {
                options.map((opt, index) => (
                    <div className="py-1 px-3 cursor-pointer hover:bg-gray-300" key={index} onClick={(e) => {
                        setText(options[index])
                        setOptions([])
                    }}>{opt}</div>
                ))
            }
            </div>}
            
        </div>
    )
}