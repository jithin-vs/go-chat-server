interface TextInputType{
  label?: string
  name: string
  type: string
  isRequired?: boolean
  placeholder: string
}
export default function TextInput({ label, name, type, isRequired = false,placeholder }: TextInputType ) {
  return (
    <>
     <label htmlFor={label} className="block text-gray-700 font-medium mb-1">
        {label}
      </label>
      <input
        // id={type}
        name={name}
        type={type}
        placeholder={placeholder}
        required={isRequired}
        className="w-full border-2 border-gray-400 rounded-lg px-4 py-2 text-gray-800 placeholder-gray-500 focus:outline-none focus:border-blue-500 mb-4"
      />
    </>
  )
}