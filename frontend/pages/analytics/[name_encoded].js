import { useRouter } from 'next/router';
import Image from 'next/image';
import axios from 'axios';
import logo from '../../public/logo.svg';
import Link from 'next/link';

export default function SenatorStats(props) {
	const router = useRouter();
	const { name_encoded } = router.query;

	const redOrGreenMoney = (number) => {
		if (number > 0) {
			return (
				<span className="text-green-500">
					{number.toLocaleString('en-US', {
						style: 'currency',
						currency: 'USD',
					})}
				</span>
			);
		} else {
			return (
				<span className="text-red-500">
					{number.toLocaleString('en-US', {
						style: 'currency',
						currency: 'USD',
					})}
				</span>
			);
		}
	};

	const redOrGreen = (number) => {
		if (number > 0) {
			return <span className="text-green-500">{number.toPrecision(4)}</span>;
		} else {
			return <span className="text-red-500">{number.toPrecision(4)}</span>;
		}
	};

  const redOrGreenPercent = (number) => {
		if (number > 0) {
			return <span className="text-green-500">{number.toPrecision(4)}</span>;
		} else {
			return <span className="text-red-500">{number.toPrecision(4)}</span>;
		}
	};

	console.log(props);

	if (!props) {
		return (
			<div className="w-screen min-h-screen flex items-center flex-col topo">
				<div className="w-1/3 pt-20 pb-20 flex flex-col">
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
				<div className="flex items-center justify-center text-4xl text-red-500 font-bold">
					We ran into an issue on our backend! Please try again.
				</div>
			</div>
		);
	}

	const name_decoded = Buffer.from(
		decodeURIComponent(name_encoded),
		'base64'
	).toString('ascii');

  const orders = props.totalOrders

	return (
		<div className="w-screen min-h-screen flex items-center flex-col topo">
			<div className="w-1/3 pt-20 pb-20 flex flex-col">
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
			<div className="pb-10 w-1/2">
				<p className="text-5xl font-bold text-red-500">{name_decoded}</p>
				<p className="text-right text-2xl font-bold py-2">
					<i>has been trading quite a bit. Do you think it's fishy?</i>
				</p>
			</div>
			<div className="w-1/2">
				{props.data &&
					Object.keys(props.data).map((key, i) => {
						const obj = props.data[key];
						return (
							<div className="mb-10" key={i}>
								<p className="text-4xl font-bold">{obj.ticker}</p>
								<div className="grid grid-cols-6 gap-4">
									<div className="col-span-2 text-center shadow-xl px-5 py-5 rounded-lg bg-white">
										<div className="text-left font-thin">POL Delta shares</div>
										<div className="text-green-500 text-5xl font-bold">
											{redOrGreen(obj.shareDelta)}
										</div>
									</div>
									<div className="col-span-2 text-center shadow-xl px-5 py-5 rounded-lg bg-white">
										<div className="text-left font-thin">
											POL Change in Value:{' '}
										</div>
										<div className="text-green-500 text-5xl font-bold">
											{redOrGreenMoney(obj.senatorValue)}
										</div>
									</div>
									<div className="col-span-2 text-center shadow-xl px-5 py-5 rounded-lg bg-white">
										<div className="text-left font-thin">
											You would have made:{' '}
										</div>
										<div className="text-green-500 text-5xl font-bold">
											{redOrGreenMoney(obj.retailValue)}
										</div>
									</div>
									<div className="col-span-2 text-center shadow-xl px-5 py-5 rounded-lg bg-white">
										<div className="text-left font-thin">
											Your Value / Senator Value:{' '}
										</div>
										<div className="text-green-500 text-5xl font-bold">
											{redOrGreen(
												obj.retailValue / obj.senatorValue * 100
											)}
										</div>
									</div>
									<div className="col-span-2 text-center shadow-xl px-5 py-5 rounded-lg bg-white">
										<div className="text-left font-thin">
											Percent of Portfolio:{' '}
										</div>
										<div className="text-green-500 text-5xl font-bold">
											{redOrGreenPercent(
												obj.totalOrders / orders * 100 
											)}
										</div>
									</div>

                  <div className="col-span-2 text-center shadow-xl px-5 py-5 rounded-lg bg-white">
										<div className="text-left font-thin">
											Peak Shares*:{' '}
										</div>
										<div className="text-green-500 text-5xl font-bold">
											{redOrGreen(
												obj.peakShares
											)}
										</div>
									</div>
								</div>
							</div>
						);
					})}
			</div>
		</div>
	);
}

export async function getServerSideProps(context) {
	const name_decoded = Buffer.from(
		decodeURIComponent(context.params.name_encoded),
		'base64'
	).toString('ascii');

	try {
		const resp = await axios.post('http://localhost:6969/shills/gabe', {
			name: name_decoded,
		});
		return {
			props: resp.data,
		};
	} catch (e) {
		return {
			props: null,
		};
	}
}
