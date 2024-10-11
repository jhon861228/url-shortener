import type React from "react";
import { useEffect, useState } from "react";
import { useCreateUrl } from "../hooks/useCreateUrl";
import { Input } from "./ui/input";
import { Button } from "./ui/button";
import CopyButton from "./CopyButton";
import { Toaster, toast } from 'react-hot-toast';
import IconQrcode from "./ui/IconDownload";
import { QRCodeSVG } from "qrcode.react";
import IconDownload from "./ui/IconDownload";


const UrlShortenCreation: React.FC = () => {
	const [urlOriginal, setUrlOriginal] = useState("");
	const { urlShorted, createUrl } = useCreateUrl();

	useEffect(() => {
		if (urlShorted) {
			console.log("URL acortada", urlShorted);
			toast.success('Tu URL fue acortada con éxito');
		}
	}, [urlShorted]);

	const handleSubmit = () => {
		createUrl(urlOriginal);
	};

	const downloadQR = (url: string) => {
		const svg = document.getElementById(`qr-code-${url}`)
		const svgData = new XMLSerializer().serializeToString(svg!)
		const canvas = document.createElement('canvas')
		const ctx = canvas.getContext('2d')
		const img = new Image()
		img.onload = () => {
		  canvas.width = img.width
		  canvas.height = img.height
		  ctx!.drawImage(img, 0, 0)
		  const pngFile = canvas.toDataURL('image/png')
		  const downloadLink = document.createElement('a')
		  downloadLink.download = `qr-code-${url}.png`
		  downloadLink.href = pngFile
		  downloadLink.click()
		}
		img.src = 'data:image/svg+xml;base64,' + btoa(svgData)
	  }

	return (
		<>
			<Toaster/>
			<div className="flex space-x-2">
				<Input
					type="url"
					id="url-input"
					placeholder="Enter your long URL here"
					onChange={(e) => setUrlOriginal(e.target.value)}
					className="flex-grow px-4 py-2 rounded-md border-2 border-teal-300 focus:border-teal-500 focus:ring focus:ring-teal-200 transition"
					required
				/>
				<Button
					onClick={() => handleSubmit()}
					className="bg-gradient-to-r from-green-500 to-blue-500 hover:from-green-600 hover:to-blue-600 text-white font-bold py-2 px-4 rounded-md transition duration-300 ease-in-out transform hover:-translate-y-1 hover:scale-105"
				>
					Acortar
				</Button>
			</div>

			{urlShorted && (
				<div className="mt-6 p-4 bg-green-100 rounded-md">
					<p className="text-sm text-gray-600 mb-2">Tu URL acortada:</p>
					<div className="flex">
						<a
							href={urlShorted.url}
							target="_blank"
							rel="noopener noreferrer"
							className="text-blue-500 hover:text-blue-700 font-medium break-all"
						>
							{urlShorted.url}
						</a>
						<CopyButton url={urlShorted.url} />
					</div>
					<div className="flex flex-col items-center">
						<QRCodeSVG id={`qr-code-${urlShorted.url}`} value={urlShorted.url} size={200} />
						<Button
							onClick={() => downloadQR(urlShorted.url)}
							className="mt-4 bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-md transition duration-300 ease-in-out flex items-center justify-center"
							>
						<IconDownload className="mr-2 h-4 w-4" />
							Download QR Code
						</Button>
					</div>
				</div>
			)}
		</>
	);
};
export default UrlShortenCreation;
