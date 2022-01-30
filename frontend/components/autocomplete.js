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

        setDebounce(setTimeout(() => {
            // TODO(dk): call api
        }, 500))
    }

    return (
        <div className="w-full flex-col">
            <input className="border-2 py-2 px-3 w-full" type="text" onInput={onType} value={text}></input>
            { options.length > 0 && <div className="relative shadow-md w-full">
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