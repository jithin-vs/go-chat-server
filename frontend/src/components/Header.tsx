export default function Header() {
  return (
    <div className="px-5 py-5 flex justify-between items-end bg-gray-600 border-b-2">
    <div className="font-semibold text-2xl">GoingChat</div>
    <div className="flex justify-between gap-3">
        <div className="">
        <input
            type="text"
            name=""
            id=""
            placeholder="search"
            className="rounded-2xl bg-gray-100 py-3 px-5 w-full"
        />
        </div>
        <div
        className="h-12 w-12 p-2 bg-blue-500 rounded-full text-white font-semibold flex items-center justify-center"
        >
        RA
        </div>
    </div>      
  </div>
  )
}