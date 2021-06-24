package bases

func GetEntityBody(domain string) string {
	return "public class " + domain + "{ \n}"
}
