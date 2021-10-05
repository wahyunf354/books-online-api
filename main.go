package main

import (
	"books_online_api/app/middlewares"
	"books_online_api/app/routes"
	_bookTypeUsecase "books_online_api/business/book_types"
	_booksUsecase "books_online_api/business/books"
	_googleBooksUsecase "books_online_api/business/google_books"
	_ordersUsecase "books_online_api/business/orders"
	_paymentMethodUsecase "books_online_api/business/payment_methods"
	_transactionsUsecase "books_online_api/business/transactions"
	_userUseCase "books_online_api/business/users"
	_bookTypeController "books_online_api/controllers/book_types"
	_booksController "books_online_api/controllers/books"
	_googleBookController "books_online_api/controllers/google_books"
	_orderController "books_online_api/controllers/orders"
	_paymentMethodController "books_online_api/controllers/payment_methods"
	_transactionsController "books_online_api/controllers/transactions"
	_userController "books_online_api/controllers/users"
	_bookDetailDb "books_online_api/drivers/databases/book_details"
	_bookDetailsDb "books_online_api/drivers/databases/book_details"
	_bookTypeDb "books_online_api/drivers/databases/book_types"
	_booksDb "books_online_api/drivers/databases/books"
	_imagesBookDb "books_online_api/drivers/databases/image_books"
	_orderDetailsDb "books_online_api/drivers/databases/order_details"
	_ordersDb "books_online_api/drivers/databases/orders"
	_paymentMethodsDb "books_online_api/drivers/databases/payment_methods"
	_transactionsDB "books_online_api/drivers/databases/transactions"
	_userDb "books_online_api/drivers/databases/users"
	_userRepository "books_online_api/drivers/databases/users"
	_booksLocal "books_online_api/drivers/localy/book_files"
	_imagesBookLocal "books_online_api/drivers/localy/image_books_files"
	_mysqlDriver "books_online_api/drivers/mysql"
	_googleBooksAPIThirtPart "books_online_api/drivers/thirdparts/google_books"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`ex.config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Server RUN on Debug mode")
	}
}

func DbMigration(db *gorm.DB) {
	var err error
	err = db.AutoMigrate(
		&_userDb.Users{},
		&_bookTypeDb.BookType{},
		&_booksDb.Book{},
		&_bookDetailsDb.BookDetails{},
		&_imagesBookDb.ImageBooks{},
		&_ordersDb.Orders{},
		&_orderDetailsDb.OrderDetails{},
		&_paymentMethodsDb.PaymentMethod{},
		&_transactionsDB.Transactions{})

	if err != nil {
		panic(err)
	}
}

func main() {
	// init connection database
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.prod.user`),
		DB_Password: viper.GetString(`database.prod.pass`),
		DB_Host:     viper.GetString(`database.prod.host`),
		DB_Port:     viper.GetString(`database.prod.port`),
		DB_Database: viper.GetString(`database.prod.name`),
	}

	configJwt := middlewares.ConfigJwt{
		SecretJwt:       viper.GetString("jwt.secret"),
		ExpiredDuration: viper.GetInt("jwt.expired"),
	}

	Conn := configDB.InitialDB()
	DbMigration(Conn)

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userRepository.NewMysqlUserRepository(Conn)
	userUseCase := _userUseCase.NewUserUseCase(userRepository, timeoutContext, configJwt)
	userController := _userController.NewUserController(userUseCase)

	googleBooksThirtPart := _googleBooksAPIThirtPart.NewGoogleBooksAPIThirtPart()
	googleBooksUsecase := _googleBooksUsecase.NewGoogleBookThirtPartUsecase(googleBooksThirtPart, timeoutContext)
	googleBooksController := _googleBookController.NewGoogleBooksController(googleBooksUsecase)

	bookTypeRepository := _bookTypeDb.NewMysqlBookTypesRepository(Conn)
	bookTypeUsecase := _bookTypeUsecase.NewBooksUseCase(bookTypeRepository, timeoutContext)
	bookTypeController := _bookTypeController.NewBookTypesController(bookTypeUsecase)

	imagesBookRepository := _imagesBookDb.NewMysqlImageBookRepository(Conn)
	imagesBookLocal := _imagesBookLocal.NewImageBookLocal(imagesBookRepository, timeoutContext)
	booksDetailRepository := _bookDetailDb.NewBookDetailsRepository(Conn, imagesBookLocal)
	booksRepository := _booksDb.NewBookRepository(Conn, booksDetailRepository)
	booksFileLocal := _booksLocal.NewBookFileLocal(booksRepository, timeoutContext)
	booksUsecae := _booksUsecase.NewBookUsecase(booksFileLocal, booksRepository, timeoutContext)
	booksController := _booksController.NewBooksController(booksUsecae)

	orderDetailsRepository := _orderDetailsDb.NewMysqlOrderDetailsRepository(Conn)
	ordersRepository := _ordersDb.NewMysqlOrderRepository(Conn, orderDetailsRepository)
	ordersUsecase := _ordersUsecase.NewOrderUsecase(ordersRepository,orderDetailsRepository, timeoutContext)
	ordersController := _orderController.NewOrdersController(ordersUsecase)

	paymentMethodRepository := _paymentMethodsDb.NewMysqlPaymentMethodsRepository(Conn)
	paymentMethodUsecase := _paymentMethodUsecase.NewPaymentMethodUsecase(paymentMethodRepository)
	paymentMethodController := _paymentMethodController.NewPaymentMethodsController(paymentMethodUsecase)

	transactionsRepository := _transactionsDB.NewTransactionsRepository(Conn, paymentMethodRepository, ordersRepository)
	transactionUsecase := _transactionsUsecase.NewTransactionUsecase(transactionsRepository)
	transactionController := _transactionsController.NewTransactionsControllers(transactionUsecase)

	routesInit := routes.ControllerList{
		JWTMiddleware:         configJwt.Init(),
		UserController:        *userController,
		GoogleBooksController: *googleBooksController,
		BookTypeController:    *bookTypeController,
		BooksController:       *booksController,
		OrdersController:      *ordersController,
		PaymentMethodsController: *paymentMethodController,
		TransactionsController: *transactionController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}
