import { classNames } from "@/utils";

export default function MessageInput(props: React.InputHTMLAttributes<HTMLInputElement>) {
    return (
        <input
          {...props}
          className={classNames(
            "block w-full rounded-xl outline outline-[1px] outline-zinc-400 border-0 py-4 px-5 bg-secondary text-black font-light placeholder:text-white/70",
            props?.className || ""
          )}
        />
      );
}