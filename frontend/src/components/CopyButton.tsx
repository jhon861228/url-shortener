import {
	Tooltip,
	TooltipContent,
	TooltipProvider,
	TooltipTrigger,
} from "@/components/ui/tooltip";
import { Button } from "@/components/ui/button";
import { useClipboard } from "@/hooks/useClipboard";

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
							<span className="h-4 w-4">Check</span>
						) : (
							<span className="h-4 w-4">Copy</span>
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
