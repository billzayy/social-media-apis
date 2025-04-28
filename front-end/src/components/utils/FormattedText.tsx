export const renderFormattedText = (input: string) => {
    const parts = input.split(/(\*\*[^*]+\*\*|\*[^*]+\*)/g); // split text by **bold** or *italic*
    return parts.map((part, index) => {
        if (part.startsWith("**") && part.endsWith("**")) { // return bold
            const cleanText = part.slice(2, -2); // remove **
            return (
                <span key={index} className="font-bold">
                    {cleanText}
                </span>
            );
        } else if (part.startsWith("*") && part.endsWith("*")) {  // return italic
            const cleanText = part.slice(1, -1); // remove *

            return (
                <span key={index} className="italic">
                    {cleanText}
                </span>
            )
        }
        return <span key={index}>{part}</span>;
    });
};