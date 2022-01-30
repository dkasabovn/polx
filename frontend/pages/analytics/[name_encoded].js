import { useRouter } from "next/router";
import Image from "next/image";
import logo from "../../public/logo.svg";

export default function SenatorStats() {
  const router = useRouter();
  const { name_encoded } = router.query;

  const name_decoded = Buffer.from(
    decodeURIComponent(name_encoded),
    "base64"
  ).toString("ascii");

  return (
    <div className="w-screen min-h-screen flex items-center flex-col topo">
      <div className="w-1/3 pt-20 pb-20 flex flex-col">
        <Image src={logo}></Image>
        <div className="flex flex-row w-full justify-around pt-5">
          <p>Home</p>
          <p>
            Politicans a.k.a <i>shills</i>
          </p>
          <p>FAQ</p>
        </div>
      </div>
      <div className="pb-10 w-1/2">
        <p className="text-5xl font-bold text-red-500">{name_decoded}</p>
        <p className="text-right text-2xl font-bold py-2">
          <i>has been trading quite a bit. Do you think it's fishy?</i>
        </p>
      </div>
      <div className="w-1/2">
        <p className="text-4xl font-bold">GOOG</p>
        <div className="grid grid-cols-6 gap-4">
          <div className="col-span-2 text-center shadow-xl px-5 py-5 rounded-lg bg-white">
            <div className="text-left font-thin">Delta shares</div>
            <div className="text-green-500 text-5xl font-bold">+4</div>
          </div>
          <div className="col-span-2 text-center shadow-xl px-5 py-5 rounded-lg bg-white">
            <div className="text-left font-thin">
              Over the past few months:{" "}
            </div>
            <div className="text-green-500 text-5xl font-bold">$100,000</div>
          </div>
          <div className="col-span-2 text-center shadow-xl px-5 py-5 rounded-lg bg-white">
            <div className="text-left font-thin">You would have made: </div>
            <div className="text-green-500 text-5xl font-bold">$100,000</div>
          </div>
        </div>
      </div>
    </div>
  );
}
