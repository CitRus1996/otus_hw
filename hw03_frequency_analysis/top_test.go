package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = false

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var ENText = `Humanity has a lot of problems. Most of them are  local
	and do not affect the global community but there are a few global
	issues. A global issue is an issue that affects every country  in 
	the world. People can clearly define them, however they are still
	not solved. One of the greatest challenges of our time is solving
	global problems.
		The United Nations is the  organization  that  defines  these
	problems and takes measures to solve them. The United Nations  is
	an example of cooperation of the  world-leading  countries.  They 
	want to make our world a better place.
		So, what global issues do they solve?  First,  it  is  global 
	warming. Our industrial activity causes a  rise  in  the  average 
	temperature. It means  that  an  average  summer  or  winter  day 
	nowadays is warmer than  it  was  a  hundred  years  ago.  Global 
	warming is  dangerous  because  it  can  cause  ice  melting.  If 
	glaciers melt, the sea level will rise, and most  coastal  cities 
	and islands will be drowned.
		Second, it is human impact on the environment. In some cities
	air pollution reached the point where it is dangerous  for  human
	health. Our plants, factories, cars produce  a  lot  of  CO2  and
	various pollutants.
		Third, overpopulation. The population of our  planet  is  7.8
	billion now, which is almost 8  times  as  much  as  it  was  two 
	hundred years ago. It will keep growing,  and  the  planet  might 
	just not be able to sustain that many inhabitants in the  future. 
	Overpopulation causes multiple problems  including  poverty.  Did 
	you know that 42 percent of people from Sub-Saharan Africa  still
	live below the poverty line? The international  poverty  line  is 
	1.9 dollars per day. People cannot fulfill their basic needs such
	as food, water and clothing.
		These are the most discussed global problems nowadays.  Their
	list might increase but all the problems  are  being  worked  on. 
	Some countries decided to switch to alternative energy sources in
	the nearest future. Most countries at least recognize the 
	problems.
		In conclusion,  I  would  like  to  say  that  not  only  the 
	governments are responsible for  our  planet,  but  every  single 
	human as well. If we want to make the world a  better  place,  we 
	should realize that even small actions affect it.`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})
}

func TestENTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		expected := []string{
			"the",    // 18
			"is",     // 12
			"of",     // 9
			"a",      // 7
			"and",    // 7
			"global", // 7
			"that",   // 7
			"to",     // 7
			"are",    // 6
			"it",     // 6
		}
		require.Equal(t, expected, Top10(ENText))
	})
}
