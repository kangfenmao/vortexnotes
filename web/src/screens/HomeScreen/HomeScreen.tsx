const HomeScreen: React.FC = () => {
  return (
    <div className='flex flex-col m-auto items-center'>
      <h1 className='text-3xl font-bold mb-10'>Vortex Notes</h1>
      <div className='relative mb-40'>
        <input
          type='text'
          placeholder='Search'
          className='p-4 py-2 w-96 border border-gray-300 rounded-md focus:outline-none focus:border-blue-500'
        />
        <button
          className='absolute top-0 right-0 h-full px-3 py-2 bg-transparent text-white rounded-r-md hover:bg-blue-600'>
          <svg
            data-v-271fb9=''
            viewBox='0 0 16 16'
            width='1em'
            height='1em'
            focusable='false'
            role='img'
            aria-label='search'
            xmlns='http://www.w3.org/2000/svg'
            fill='currentColor'>
            <g data-v-271fb9=''>
              <path
                d='M11.742 10.344a6.5 6.5 0 1 0-1.397 1.398h-.001c.03.04.062.078.098.115l3.85 3.85a1 1 0 0 0 1.415-1.414l-3.85-3.85a1.007 1.007 0 0 0-.115-.1zM12 6.5a5.5 5.5 0 1 1-11 0 5.5 5.5 0 0 1 11 0z'></path>
            </g>
          </svg>
        </button>
      </div>
    </div>
  )
}

export default HomeScreen
