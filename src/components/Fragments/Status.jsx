const Status = () => {
  return (
    <div className="w-full h-full bg-white border-gray-700 rounded-lg shadow my-2 flex flex-col justify-between">
      <div class="flex flex-col">
        <div class="overflow-x-auto sm:-mx-6 lg:-mx-8">
          <div class="inline-block min-w-full py-2 sm:px-6 lg:px-8">
            <div class="overflow-hidden">
              <div className="h-64 bg-blue-700 border-gray-700 rounded-lg shadow mt-12 mx-5 justify-between">
                <p className="text-4xl font-bold text-white w-full h-full flex flex-col h-full w-full justify-center items-center">
                  DOPIN
                </p>
              </div>
            </div>
          </div>
        </div>

        <div class="overflow-x-auto sm:-mx-6 lg:-mx-8">
          <div class="inline-block min-w-full py-2 sm:px-6 lg:px-8">
            <div class="overflow-hidden">
              <div className="justify-between w-full">
                <div className="mx-16 bg-yellow-300 border-gray-700 rounded-lg shadow py-1">
                  <p className="text-3xl font-bold text-black w-full h-full flex flex-col h-full w-full justify-center items-center">
                    Menunggu Pemindaian Uang...
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Status;
