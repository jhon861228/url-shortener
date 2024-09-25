import type React from "react";
import { useState } from "react";
import { useCreateUrl } from "../hooks/useCreateUrl";
import { Input } from "./ui/input";
import { Button } from "./ui/button";
import CopyButton from "./CopyButton";

const UrlShortenCreation: React.FC = () => {
	const [urlOriginal, setUrlOriginal] = useState("");

	const { urlShorted, createUrl } = useCreateUrl();

	const handleSubmit = () => {
		createUrl(urlOriginal);
	};

	return (
		<>
			<div className="space-y-4">
				<div className="space-y-2">
					<Input
						type="url"
						id="url-input"
						placeholder="Enter your long URL here"
						onChange={(e) => setUrlOriginal(e.target.value)}
						className="w-full px-4 py-2 rounded-md border-2 border-teal-300 focus:border-teal-500 focus:ring focus:ring-teal-200 transition"
						required
					/>
				</div>
				<Button
					onClick={() => handleSubmit()}
					className="w-full bg-gradient-to-r from-green-500 to-blue-500 hover:from-green-600 hover:to-blue-600 text-white font-bold py-2 px-4 rounded-md transition duration-300 ease-in-out transform hover:-translate-y-1 hover:scale-105"
				>
					Acortar
				</Button>
			</div>

			{urlShorted && (
				<div className="mt-6 p-4 bg-green-100 rounded-md">
					<p className="text-sm text-gray-600 mb-2">Tu URL acortada:</p>
					<div className="flex">
						<a
							href={urlShorted}
							target="_blank"
							rel="noopener noreferrer"
							className="text-blue-500 hover:text-blue-700 font-medium break-all"
						>
							{urlShorted}
						</a>
						<CopyButton url={urlShorted} />
					</div>
				</div>
			)}
		</>
	);
};
export default UrlShortenCreation;
