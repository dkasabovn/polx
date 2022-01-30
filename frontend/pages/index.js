import Head from "next/head";
import Image from "next/image";
import Autocomplete from "../components/autocomplete";
import logo from "../public/logo.svg";
import Link from "next/link";

export default function Home() {
  return (
    <div className="w-screen min-h-screen flex items-center flex-col topo">
      <div className="w-1/3 pt-20 flex flex-col">
        <Image src={logo}></Image>
        <div className="flex flex-row w-full justify-around pt-5">
          <Link href="/">
            <a>Home</a>
          </Link>
          <Link href="/shills">
            <a>
              Politicans a.k.a <i>shills</i>
            </a>
          </Link>
          <Link href="/faq">
            <a>FAQ</a>
          </Link>
        </div>
      </div>
      <div className="w-1/2 px-10 pt-20">
        <Autocomplete></Autocomplete>
      </div>
    </div>
  );
}
