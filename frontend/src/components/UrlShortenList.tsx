import react from "react";
import { useStore } from "@nanostores/react";
import CopyButton from "./CopyButton";
import { urlCreated } from "../store/urlStore";
import { useEffect, useState } from "react";
import { Button } from "./ui/button";
import { QRCodeSVG } from "qrcode.react";
import IconDownload from "./ui/IconDownload";
import {
	Collapsible,
	CollapsibleContent,
	CollapsibleTrigger,
  } from "@/components/ui/collapsible"
import type { UrlShortenModel } from "@/models/UrlShorten.model";
import { parseDate } from "@/utils/DateUtil";

const UrlShortenList: React.FC = () => {
	const [urls, setUrls] = useState<UrlShortenModel[]>([]);
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
			{urls.length > 0 && (
				<div className="mt-8">
					<h3 className="text-lg font-semibold text-gray-700 mb-4">
						Urls recientes
					</h3>
					<div className="grid grid-cols-3 gap-4">
                {displayedHistory.map((entry, index) => (
                  <div key={index} className="bg-teal-50 p-3 rounded-md">
                    <div className="flex flex-col h-full">
                      <div className="flex-grow">
                        <a
                          href={entry.url}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="text-blue-500 hover:text-blue-700 font-medium break-all"
                        >
                          {entry.url}
                        </a>
                        <p className="text-xs text-gray-500 mt-1">
                          {entry.createdAt}
                        </p>
                      </div>
                      <div className="flex items-center justify-between mt-2">
                        <CopyButton url={entry.url} />
                        <Collapsible>
                          <CollapsibleTrigger asChild>
                            <Button variant="ghost" size="icon">
                              <QRCodeSVG className="h-4 w-4" />
                              <span className="sr-only">Toggle QR Code</span>
                            </Button>
                          </CollapsibleTrigger>
                          <CollapsibleContent className="mt-2">
                            <div className="flex flex-col items-center">
                              <QRCodeSVG id={`qr-code-${entry.url}`} value={entry.url} size={100} />
                              <Button
                                className="mt-2 bg-blue-500 hover:bg-blue-600 text-white font-bold py-1 px-2 rounded-md transition duration-300 ease-in-out flex items-center justify-center text-sm"
                              >
                                <IconDownload className="mr-1 h-3 w-3" />
                                Download QR
                              </Button>
                            </div>
                          </CollapsibleContent>
                        </Collapsible>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
					{urls.length > 3 && (
						<Button
							onClick={toggleHistory}
							className="mt-4 w-full bg-teal-100 text-teal-700 hover:bg-teal-200 font-medium py-2 px-4 rounded-md transition duration-300 ease-in-out flex items-center justify-center"
						>
							{showAllHistory ? (
								<>
									Mostrar menos <span className="ml-2 h-4 w-4">-</span>
								</>
							) : (
								<>
									Mostrar mas <span className="ml-2 h-4 w-4">+</span>
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
