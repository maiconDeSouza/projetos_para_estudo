package name

import (
	"fmt"
	"igressos-goroutines/num"
)

var firstName = []string{
	"Abigail", "Abner", "Adalberto", "Adauto", "Adelson", "Ademir", "Adriana", "Adélia", "Afonso", "Agnaldo",
	"Agnes", "Aguinaldo", "Aida", "Ailton", "Airton", "Alan", "Alana", "Alba", "Alcina", "Alda",
	"Alencar", "Alessandra", "Alex", "Alexandre", "Alfeu", "Alfredo", "Alice", "Aline", "Alisson", "Almir",
	"Aloísio", "Altair", "Alvaro", "Alvimar", "Amadeu", "Amanda", "Amarildo", "Amauri", "Americo", "Amilton",
	"Amir", "Amira", "Amália", "Amélia", "Ana", "Analu", "Ananias", "Anchieta", "Anderson", "Andressa",
	"André", "Andréia", "Anelise", "Angelina", "Angelo", "Angélica", "Anita", "Antônia", "Antônio", "Araci",
	"Ari", "Ariane", "Ariela", "Arlene", "Arnaldo", "Arthur", "Ary", "Asaf", "Assis", "Astrid",
	"Audrey", "Augustinho", "Augusto", "Aura", "Aurelia", "Aurora", "Avani", "Avelino", "Azarias", "Baltazar",
	"Bartolomeu", "Basílio", "Beatriz", "Bela", "Belmiro", "Benedito", "Benjamim", "Benjamin", "Bento", "Benício",
	"Berenice", "Bernadete", "Bernardo", "Berta", "Betina", "Betânia", "Bianca", "Branca", "Brena", "Breno",
	"Brian", "Briseis", "Bruce", "Bruna", "Bruno", "Bryan", "Brás", "Brígida", "Bárbara", "Cacilda",
	"Caio", "Caitlin", "Caleb", "Calinda", "Calisto", "Camila", "Camilo", "Candice", "Candido", "Carina",
	"Carla", "Carlos", "Carmem", "Carolina", "Casemiro", "Cassiano", "Catarina", "Cecília", "Celina", "Celso",
	"Cesar", "Charline", "Chayenne", "Chico", "Chloe", "Christian", "Cibele", "Cicero", "Cid", "Cidália",
	"Cildo", "Cirilo", "Ciro", "Clara", "Clarice", "Clarissa", "Claude", "Claudemir", "Claudete", "Claudinei",
	"Cleide", "Cleonice", "Clodoaldo", "Cláudia", "Cláudio", "Cléo", "Clóvis", "Constança", "Cora", "Cordélia",
	"Cosme", "Cristiana", "Cristiano", "Cristina", "Custódia", "Cynthia", "Cássio", "Célia", "César", "Cíntia",
	"Dafne", "Dagoberto", "Dalila", "Dalmo", "Dalton", "Dalva", "Damaris", "Damiao", "Dana", "Daniel",
	"Daniela", "Danilo", "Darci", "Darcio", "Darius", "Davi", "Deivid", "Delfim", "Demétrio", "Dener",
	"Denilson", "Denis", "Denzel", "Derick", "Diana", "Diego", "Djalma", "Djalmo", "Dom", "Domingos",
	"Donizete", "Dorival", "Douglas", "Duarte", "Durval", "Dácio", "Dália", "Débora", "Edemilson", "Eden",
	"Edenilson", "Eder", "Ederson", "Edgard", "Edilson", "Edmar", "Edmilson", "Edmo", "Edmundo", "Edson",
	"Eduardo", "Edward", "Edwin", "Efrain", "Egberto", "Egon", "Elaine", "Elcio", "Eliaquim", "Elias",
	"Eliel", "Eliezer", "Elinei", "Elisa", "Eliseu", "Elismar", "Elmo", "Elpídio", "Elson", "Elton",
	"Emanuel", "Enzo", "Erica", "Erick", "Fabiana", "Fabrício", "Felipe", "Fernanda", "Fernando",
}

var lastNames = []string{
	"Abrantes", "Abreu", "Aguiar", "Aires", "Albuquerque", "Alcantara", "Aldeia", "Aleixo", "Alencar", "Almeida",
	"Alonço", "Alvarenga", "Alvares", "Alves", "Amado", "Amaral", "Amorim", "Anchieta", "Andrada", "Andrade",
	"Antunes", "Anunciação", "Aranha", "Araújo", "Arce", "Arruda", "Assis", "Assunção", "Atalaia", "Avelar",
	"Avila", "Azeredo", "Azevedo", "Bacelar", "Bahia", "Balduino", "Baptista", "Barata", "Barbosa", "Barreiros",
	"Barreto", "Barros", "Barroso", "Bastos", "Batista", "Belfort", "Belém", "Bemerguy", "Benites", "Bentes",
	"Bermudes", "Bernandes", "Bernardes", "Bicalho", "Bicudo", "Bissoli", "Bitencourt", "Blanco", "Boaventura", "Bogado",
	"Borges", "Borja", "Botelho", "Braga", "Braganca", "Bragança", "Branco", "Brandão", "Brasileiro", "Brito",
	"Brizola", "Brochado", "Brum", "Bueno", "Bugre", "Bulhões", "Cabral", "Cabrera", "Caetano", "Caires",
	"Calado", "Calazans", "Caldeira", "Caldas", "Caldeiras", "Camacho", "Camargo", "Caminha", "Camões", "Campelo",
	"Campos", "Canário", "Capanema", "Capucho", "Cardoso", "Carmo", "Carneiro", "Carreira", "Carreiro", "Carvalho",
	"Casado", "Cascais", "Casimiro", "Castelo", "Castilho", "Castro", "Catanhede", "Catarino", "Cavalcante", "Cavalcanti",
	"Cazumba", "Cerejeira", "Cerveira", "Chagas", "Chaves", "Cintra", "Cipriano", "Clemente", "Coelho", "Coimbra",
	"Colares", "Conceição", "Conde", "Cordeiro", "Correa", "Correia", "Corta", "Cortez", "Costa", "Coutinho",
	"Couto", "Covas", "Crespo", "Cruz", "Cunha", "Damasco", "Dantas", "Delfim", "Delgado", "Dias",
	"Diniz", "Diorio", "Domingues", "Donato", "Dorneles", "Dornelles", "Dourado", "Duarte", "Duraes", "Durães",
	"Echeverria", "Elias", "Encarnação", "Engler", "Espinola", "Espinosa", "Esteves", "Evangelista", "Fagundes", "Faria",
	"Farias", "Faro", "Faustino", "Feijó", "Feitosa", "Feliciano", "Felix", "Fernandes", "Ferrão", "Ferraz",
	"Ferreira", "Ferro", "Fialho", "Figueira", "Figueiredo", "Filgueiras", "Firme", "Fiuza", "Florentino", "Flores",
	"Fogaça", "Fonseca", "Fontes", "Forjaz", "Fortes", "Frade", "Fraga", "Fragoso", "França", "Franco",
	"Freire", "Freitas", "Frota", "Furtado", "Gabriel", "Gago", "Galvão", "Gama", "Gamelas", "Garcês",
	"Garcia", "Garrido", "Gaspar", "Gato", "Gatto", "Geraldes", "Gil", "Gimenez", "Giroto", "Godoi",
	"Godinho", "Godoy", "Goes", "Góis", "Gomes", "Gonçalves", "Gordilho", "Goulart", "Gouveia", "Gouveiaa",
}

func GenerateNames(v int) []string {
	var names []string

	for i := 1; i <= v; i++ {
		first := firstName[num.GenerateNumberRand(0, len(firstName)-1)]
		last := lastNames[num.GenerateNumberRand(0, len(lastNames)-1)]
		names = append(names, fmt.Sprintf("%s %s", first, last))
	}

	return names
}
