export default function SubmitButton({ Text }: { Text: string }) {
  return (
    <div className="flex justify-center">
      <button
        type="submit"
        className="w-full bg-black text-white font-semibold rounded-3xl py-3 mt-4 "
      >
       {Text}
      </button>
  </div>
  )
}