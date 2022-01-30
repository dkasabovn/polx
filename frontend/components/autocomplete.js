import { useRouter } from "next/router"
import { useState } from "react"

export default function Autocomplete() {
    const router = useRouter()
    const [text, setText] = useState("")
    const [options, setOptions] = useState([])
    const [debounce, setDebounce] = useState(null)

    const onType = (e) => {
        setText(e.target.value)

        if (debounce) {
            clearTimeout(debounce)
        }

        setDebounce(setTimeout(async () => {
            fetch(`/api/autocomplete?value=${text}`).then((x) => {
                x.json().then((opts) => {
                    if (opts.data) {
                        const res_data = opts.data.map((x) => {
                            return x.name
                        })
                        setOptions([...res_data])
                    }
                }).catch((err) => { console.log(err) })
            }).catch((err) => { console.log(err) })
        }, 100))
    }

    const onSubmitSearch = (e) => {
        router.push(`/analytics/${encodeURIComponent(Buffer.from(text).toString("base64"))}`)
    }

    return (
        <div className="w-full flex-col">
            <div className="w-full flex-row">
            <input className="border-l-2 border-t-2 border-b-2 py-2 px-3 w-5/6" type="text" onInput={onType} onKeyDown={(e) => {
                if (e.key === 'Enter') {
                    onSubmitSearch(e)
                }
            }} value={text}></input>
            <button className="w-1/6 border-r-2 border-t-2 border-b-2 py-2 border-l-slate-100 border-l-2 bg-white" onClick={onSubmitSearch}>Search</button>
            </div>
            { options.length > 0 && <div className="relative shadow-md w-full">
            {
                options.map((opt, index) => (
                    <div className="py-1 px-3 cursor-pointer hover:bg-gray-300 bg-white" key={index} onClick={(e) => {
                        setText(options[index])
                        setOptions([])
                    }}>{opt}</div>
                ))
            }
            </div>}
            
        </div>
    )
}