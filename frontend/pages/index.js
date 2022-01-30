import Head from "next/head";
import Image from "next/image";
import Autocomplete from "../components/autocomplete";
import logo from "../public/logo.svg";

export default function Home() {
  return (
    <div className="w-screen min-h-screen flex items-center flex-col">
      <div className="w-1/3 pt-20 flex flex-col">
        <Image src={logo}></Image>
        <div className="flex flex-row w-full justify-around pt-5">
          <p>Home</p>
          <p>
            Politicans a.k.a <i>shills</i>
          </p>
          <p>FAQ</p>
        </div>
      </div>
      <div className="w-1/2 px-10 pt-20">
        <Autocomplete></Autocomplete>
      </div>
    </div>
  );
}
