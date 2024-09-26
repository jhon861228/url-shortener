import react from "react";
import { useStore } from "@nanostores/react";
import CopyButton from "./CopyButton";
import { urlCreated } from "../store/urlStore";
import { useEffect, useState } from "react";
import { message } from "antd";
import { Button } from "./ui/button";
import { ChevronDown, ChevronUp } from "lucide-react";

const UrlShortenList: React.FC = () => {
	const [urls, setUrls] = useState([]);
	const [contextHolder] = message.useMessage();
	const [showAllHistory, setShowAllHistory] = useState(false);

	const $urlStore = useStore(urlCreated);

	useEffect(() => {
		setUrls(JSON.parse(window.localStorage.getItem("urlShorten") || "[]"));
	}, [$urlStore]);

	const toggleHistory = () => {
		setShowAllHistory(!showAllHistory);
	};

	const displayedHistory = showAllHistory ? urls : urls.slice(0, 3);

	return (
		<>
			{contextHolder}
			{urls.length > 0 && (
				<div className="mt-8">
					<h3 className="text-lg font-semibold text-gray-700 mb-4">
						Urls recientes
					</h3>
					<ul className="space-y-3">
						{displayedHistory.map((entry, index) => (
							<li key={index} className="bg-teal-50 p-3 rounded-md">
								<div className="flex">
									<a
										href={entry}
										target="_blank"
										rel="noopener noreferrer"
										className="text-blue-500 hover:text-blue-700 font-medium break-all"
									>
										{entry}
									</a>
									<CopyButton url={entry} />
								</div>
							</li>
						))}
					</ul>
					{urls.length > 3 && (
						<Button
							onClick={toggleHistory}
							className="mt-4 w-full bg-teal-100 text-teal-700 hover:bg-teal-200 font-medium py-2 px-4 rounded-md transition duration-300 ease-in-out flex items-center justify-center"
						>
							{showAllHistory ? (
								<>
									Mostrar menos <ChevronUp className="ml-2 h-4 w-4" />
								</>
							) : (
								<>
									Mostrar mas <ChevronDown className="ml-2 h-4 w-4" />
								</>
							)}
						</Button>
					)}
				</div>
			)}
		</>
	);
};

export default UrlShortenList;
