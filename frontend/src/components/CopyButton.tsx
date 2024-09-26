import {
	Tooltip,
	TooltipContent,
	TooltipProvider,
	TooltipTrigger,
} from "@/components/ui/tooltip";
import { Button } from "@/components/ui/button";
import { useClipboard } from "@/hooks/useClipboard";
import { Check, Copy } from "lucide-react";

interface CopyButtonProps {
	url: string;
}

const CopyButton = ({ url }: CopyButtonProps) => {
	const { isCopied, copyToClipboard } = useClipboard();

	return (
		<TooltipProvider>
			<Tooltip>
				<TooltipTrigger asChild>
					<Button
						variant="ghost"
						size="icon"
						onClick={() => copyToClipboard(url)}
						className="ml-2 h-8 text-gray-500 hover:text-gray-700"
					>
						{isCopied ? (
							<Check className="h-4 w-4" />
						) : (
							<Copy className="h-4 w-4" />
						)}
						<span className="sr-only">Copy URL</span>
					</Button>
				</TooltipTrigger>
				<TooltipContent>
					<p>{isCopied ? "Copied!" : "Copy URL"}</p>
				</TooltipContent>
			</Tooltip>
		</TooltipProvider>
	);
};

export default CopyButton;
