const Topbar = () => {
  return (
    <div className="w-full h-fit bg-white border-gray-700 rounded-lg shadow my-2 flex flex-col justify-between">
      <div className="flex flex-row">
        <div className="flex-none w-1/2">
          <div class="overflow-x-auto sm:-mx-6 lg:-mx-8">
            <div class="inline-block min-w-full py-2 sm:px-6 lg:px-8">
              <div class="overflow-hidden">
                <p className="text-3xl font-bold ml-5 my-3">DOPIN</p>
              </div>
            </div>
          </div>
        </div>
        <div className="align-center w-full items-center">
          <p className="text-3xl font-bold py-2 ml-5 my-3">Dashboard Kasir</p>
        </div>
      </div>
    </div>
  );
};

export default Topbar;
